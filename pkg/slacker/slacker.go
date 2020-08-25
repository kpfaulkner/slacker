package slacker

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/kpfaulkner/goui/pkg"
	"github.com/kpfaulkner/goui/pkg/events"
	"github.com/kpfaulkner/goui/pkg/widgets"
	log "github.com/sirupsen/logrus"
	"image/color"
	"os"
	"time"
)

type Slacker struct {
	window pkg.Window

	userDetails  UserDetails
	contacts     []Contact
	channels     []Channel
	slackHandler *SlackHandler

	// messages to display.
	messages []Message

	// contacts/channel vpanel.

	// Will want to have a reference to these for
	// populating later.
	contactsChannelsVPanel *widgets.VPanel
	messagesVPanel         *widgets.VPanel
}

func NewSlacker() *Slacker {
	s := Slacker{}

	// just grab slack key from env var
	slackKey := os.Getenv("SLACK_KEY")
	s.slackHandler = NewSlackHandler(slackKey)

	// just keep it simple at 800x600 for now. :)
	s.window = pkg.NewWindow(800, 600, "Slacker", false, false)

	return &s
}

// ListenForIncomingMessages will loop forever listening for messages.
func (s *Slacker) ListenForIncomingMessages() error {

	for {

		// sleep a bit so we dont waste resources.
		time.Sleep(500 * time.Millisecond)
	}

	return nil
}

func (s *Slacker) QuitSlacker(event events.IEvent) error {
	// TODO(kpfaulkner)
	// Close connection to slack gracefully... then simply bomb out.

	log.Fatalf("QuitSlacker NOT IMPLEMENTED!!!\n")
	return nil
}

// SetupUI generates the UI structure of the application.
// For the purposes of this test app, we're just sticking with a fixed 800x600 window size
// which means all internal panels/widgets are are fixed known size. As more dynamic features
// are added to GoUI, this will be rewritten to take advantage of it.
func (s *Slacker) SetupUI() error {

	// black toolbars are cool.
	tb := widgets.NewToolBar("toolbar", &color.RGBA{0, 0, 0, 0xff})
	quitButton := widgets.NewToolbarItem("quitTBI", s.QuitSlacker)
	tb.AddToolBarItem(quitButton)
	tb.SetSize(800, 30)

	// Main VPanel of the app. Will have 2 entries in it. The first is the top level toolbar
	// secondly will be a HPanel to have contacts and messages.
	vpanel := widgets.NewVPanelWithSize("main-vpanel", 800, 600, &color.RGBA{0, 0, 0, 0xff})
	vpanel.AddWidget(tb)
	s.window.AddPanel(vpanel)

	// main Horizontal panel... first element is channels + people
	// second element is actual chat.
	mainHPanel := widgets.NewHPanel("hpanel1", &color.RGBA{0, 100, 0, 255})

	// 2 main sections added to mainHPanel
	// contactsVPanel goes down complete left side, 100 width, 600-30 (toolbar) in height
	s.contactsChannelsVPanel = widgets.NewVPanelWithSize("contactsVPanel", 150, 570, &color.RGBA{0, 0, 100, 0xff})

	// In messagesTypingVPanel we will have 2 vpanels.
	messagesTypingVPanel := widgets.NewVPanelWithSize("messagesTypingVPanel", 650, 570, &color.RGBA{0, 50, 50, 0xff})

	// The first for messages the second for typing widget.
	s.messagesVPanel = widgets.NewVPanelWithSize("messagesVPanel", 650, 540, &color.RGBA{10, 50, 50, 0xff})
	typingVPanel := widgets.NewVPanelWithSize("typingVPanel", 650, 30, &color.RGBA{50, 50, 50, 0xff})
	messagesTypingVPanel.AddWidget(s.messagesVPanel)
	messagesTypingVPanel.AddWidget(typingVPanel)

	mainHPanel.AddWidget(s.contactsChannelsVPanel)
	mainHPanel.AddWidget(messagesTypingVPanel)

	// now add mainHPanel to VPanel.
	vpanel.AddWidget(mainHPanel)

	return nil
}

func (s *Slacker) contactSelected(event events.IEvent) error {
	log.Debugf("Contact %s selected", event.WidgetID())

	msgList, err := s.slackHandler.GetMessagesForContact(event.WidgetID())
	if err != nil {
		return err
	}

	s.messages = msgList

	s.populateMessagesUI()
	return nil
}

func (s *Slacker) channelSelected(event events.IEvent) error {
	log.Debugf("Channel %s selected", event.WidgetID())
	msgList, err := s.slackHandler.GetMessagesForChannel(event.WidgetID())
	if err != nil {
		return err
	}
	s.messages = msgList
	s.populateMessagesUI()
	return nil
}

// populateMessagesUI takes the messages in s.Messages and creates the UI elements
// to display.
func (s *Slacker) populateMessagesUI() error {
	s.messagesVPanel.ClearWidgets()

	// messages should be...... labels?
	for i, msg := range s.messages {
		l := widgets.NewLabel(fmt.Sprintf("label-%d", i), msg.Text, 400, 40, &color.RGBA{0, 0, 0, 0xff}, nil)
		s.messagesVPanel.AddWidget(l)
	}
	return nil
}

// populateContactChannelUI is used to populate the contactsChannelsVPanel with the information in
// the contacts and channel slices
func (s *Slacker) populateContactChannelUI() error {

	for _, contact := range s.contacts {
		tb := widgets.NewTextButton(fmt.Sprintf("CT: %s button", contact.Name), "CT: "+contact.Name, true, 0, 0, nil, nil, nil, s.contactSelected)
		s.contactsChannelsVPanel.AddWidget(tb)
	}

	for _, ch := range s.channels {
		tb := widgets.NewTextButton(fmt.Sprintf("CH: %s button", ch.Name), "CH: "+ch.Name, true, 0, 0, nil, nil, nil, s.channelSelected)
		s.contactsChannelsVPanel.AddWidget(tb)
	}
	return nil
}

func (s *Slacker) generateDummyData() {
	s.channels = []Channel{Channel{Name: "channel1"}, Channel{Name: "channel2"}, Channel{Name: "channel3"}}
	s.contacts = []Contact{Contact{Name: "contact1"}, Contact{Name: "contact2"}, Contact{Name: "contact3"}}
}

func (s *Slacker) Run() {
	log.SetLevel(log.DebugLevel)
	s.SetupUI()
	ebiten.SetRunnableInBackground(true)
	ebiten.SetWindowResizable(true)

	// dummy data for now.
	s.generateDummyData()

	// generate UI with above dummy data
	s.populateContactChannelUI()

	go s.ListenForIncomingMessages()

	// UI in main loop
	s.window.MainLoop()
}
