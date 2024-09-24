package ui

func Text(text string) *TextComponent {
	return &TextComponent{Text: text}
}

func InputBox(max_size uint) *InputBoxComponent {
	return &InputBoxComponent{max_size: max_size}
}
