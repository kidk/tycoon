package ui

import (
	"image/color"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kidk/tycoon/engine"
	"github.com/kidk/tycoon/graphics"
	"github.com/kidk/tycoon/helpers"
)

type UITestScreen struct {
	ui *ebitenui.UI

	playerRenderer graphics.NPCRenderer
}

func NewUITestScreen(spriteCache *graphics.SpriteCache, mouseListener func(event *widget.WidgetMouseButtonPressedEventArgs), mouse *helpers.MouseHelper) *UITestScreen {
	// load images for button states: idle, hover, and pressed
	buttonImage, _ := LoadButtonImage()

	// load button text font
	face, _ := LoadFont(20)

	// construct a new container that serves as the root of the UI hierarchy
	rootContainer := widget.NewContainer(
		// the container will use an anchor layout to layout its single child widget
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			// Define number of columns in the grid
			widget.GridLayoutOpts.Columns(1),
			// Define how to stretch the rows and columns. Note it is required to
			// specify the Stretch for each row and column.
			widget.GridLayoutOpts.Stretch([]bool{true, false, true, true}, []bool{false, true, false, false}),
		)),
	)

	/*
		Header container
	*/
	headerContainer := widget.NewContainer(
		// the container will use a plain color as its background
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0x13, 0x1a, 0x22, 0xff})),

		// the container will use an anchor layout to lay out its single child widget
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Spacing(10),
			widget.RowLayoutOpts.Padding(widget.Insets{
				Top:    10,
				Bottom: 10,
				Left:   10,
				Right:  10,
			}),
		)),
	)
	rootContainer.AddChild(headerContainer)

	titleLabel := widget.NewText(
		widget.TextOpts.Text("Hospitality", face, color.White),
	)
	headerContainer.AddChild(titleLabel)
	moneyLabel := widget.NewText(
		widget.TextOpts.Text("$ 10000", face, color.White),
	)
	headerContainer.AddChild(moneyLabel)

	/*
		Middle container (nothing)
	*/
	middleContainer := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MouseButtonPressedHandler(mouseListener),
		),
	)
	rootContainer.AddChild(middleContainer)

	// buildingContainer := widget.NewContainer(
	// 	// the container will use a plain color as its background
	// 	widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0x13, 0x1a, 0x22, 0xff})),
	// 	// the container will use an anchor layout to layout its single child widget
	// 	widget.ContainerOpts.Layout(widget.NewRowLayout(
	// 		widget.RowLayoutOpts.Spacing(10),
	// 		widget.RowLayoutOpts.Padding(widget.Insets{
	// 			Top:    10,
	// 			Bottom: 10,
	// 			Left:   10,
	// 			Right:  10,
	// 		}),
	// 	)),
	// )
	// buildingContainer.GetWidget().Visibility = widget.Visibility_Hide
	// rootContainer.AddChild(buildingContainer)
	//
	// buildingContainer.AddChild(CreateButton("Wall", func(args *widget.ButtonClickedEventArgs) {
	// 	println("wall button clicked")
	// }, buttonImage, face))

	// buildingContainer.AddChild(CreateButton("Floor", func(args *widget.ButtonClickedEventArgs) {
	// 	println("floor button clicked")
	// }, buttonImage, face))

	// buildingContainer.AddChild(CreateButton("Door", func(args *widget.ButtonClickedEventArgs) {
	// 	println("door button clicked")
	// }, buttonImage, face))

	/*
		Footer
	*/
	footerContainer := widget.NewContainer(
		// the container will use a plain color as its background
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0x13, 0x1a, 0x22, 0xff})),

		// the container will use an anchor layout to layout its single child widget
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Spacing(10),
			widget.RowLayoutOpts.Padding(widget.Insets{
				Top:    10,
				Bottom: 10,
				Left:   10,
				Right:  10,
			}),
		)),
	)
	rootContainer.AddChild(footerContainer)

	// construct a button
	footerContainer.AddChild(CreateButton("Wall", func(args *widget.ButtonClickedEventArgs) {
		println("wall clicked")
		//buildingContainer.GetWidget().Hidden = !buildingContainer.GetWidget().Hidden
		mouse.SetCursor("building", "wall_brown_up_middle")
	}, buttonImage, face))

	footerContainer.AddChild(CreateButton("Floor", func(args *widget.ButtonClickedEventArgs) {
		println("floor clicked")
		mouse.SetCursor("floor", "floor_wood_light")
	}, buttonImage, face))

	footerContainer.AddChild(CreateButton("Reception zone", func(args *widget.ButtonClickedEventArgs) {
		println("reception zone")
		mouse.SetCursor("zone", "reception")
	}, buttonImage, face))

	footerContainer.AddChild(CreateButton("Reception desk", func(args *widget.ButtonClickedEventArgs) {
		println("reception desk")
		mouse.SetCursor("item", "reception_desk")
	}, buttonImage, face))

	footerContainer.AddChild(CreateButton("Room zone", func(args *widget.ButtonClickedEventArgs) {
		println("room zone")
		mouse.SetCursor("zone", "room")
	}, buttonImage, face))

	footerContainer.AddChild(CreateButton("Bed", func(args *widget.ButtonClickedEventArgs) {
		println("bed")
		mouse.SetCursor("item", "bed")
	}, buttonImage, face))

	// footerContainer.AddChild(CreateButton("Testing", func(args *widget.ButtonClickedEventArgs) {
	// 	println("testing")
	// 	if buildingContainer.GetWidget().Visibility == widget.Visibility_Show {
	// 		buildingContainer.GetWidget().Visibility = widget.Visibility_Hide
	// 	} else {
	// 		buildingContainer.GetWidget().Visibility = widget.Visibility_Show
	// 	}
	// }, buttonImage, face))

	// construct the UI
	ui := ebitenui.UI{
		Container: rootContainer,
	}

	player := engine.NewPlayer()

	return &UITestScreen{
		ui:             &ui,
		playerRenderer: graphics.NewNPCRenderer(spriteCache, player, 0, 0),
	}
}

func (tds *UITestScreen) Update() error {
	tds.ui.Update()

	return nil
}

func (tds *UITestScreen) Draw(screen *ebiten.Image) {
	tds.ui.Draw(screen)
}
