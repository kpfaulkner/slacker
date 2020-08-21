package slacker

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kpfaulkner/goui/pkg"
	"github.com/kpfaulkner/goui/pkg/events"
	"github.com/kpfaulkner/goui/pkg/widgets"
	log "github.com/sirupsen/logrus"
	"image/color"
)

type Slacker struct {
	window pkg.Window
}

func NewSlacker() *Slacker {
	s := Slacker{}

	// just keep it simple at 800x600 for now. :)
	s.window = pkg.NewWindow(800, 600, "Slacker", false, false)
	return &s
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
	contactsVPanel := widgets.NewVPanelWithSize("contactsVPanel", 100, 570, &color.RGBA{0, 0, 100, 0xff})

	// In messagesTypingVPanel we will have 2 vpanels.
	messagesTypingVPanel := widgets.NewVPanelWithSize("messagesTypingVPanel", 700, 570, &color.RGBA{0, 50, 50, 0xff})

	// The first for messages the second for typing widget.
	messagesVPanel := widgets.NewVPanelWithSize("messagesVPanel", 700, 540, &color.RGBA{10, 50, 50, 0xff})
	typingVPanel := widgets.NewVPanelWithSize("typingVPanel", 700, 30, &color.RGBA{50, 50, 50, 0xff})
	messagesTypingVPanel.AddWidget(messagesVPanel)
	messagesTypingVPanel.AddWidget(typingVPanel)

	mainHPanel.AddWidget(contactsVPanel)
	mainHPanel.AddWidget(messagesTypingVPanel)

	// now add mainHPanel to VPanel.
	vpanel.AddWidget(mainHPanel)

	return nil
}

func (s *Slacker) Run() {
	s.SetupUI()

	ebiten.SetRunnableInBackground(true)
	ebiten.SetWindowResizable(true)
	s.window.MainLoop()
}
