package buttons

import (
	"conver"
	"fmt"
	"gopkg.in/tucnak/telebot.v2"
	"strconv"
)

type Buttons struct {
	Button telebot.ReplyMarkup
}

func NewButtons(b telebot.ReplyMarkup) *Buttons {
	b.ResizeReplyKeyboard = true
	return &Buttons{Button: b}
}

func (s *Buttons) Main() telebot.ReplyMarkup {
	main := s.Button
	b1 := main.Text("Конвертация")
	b2 := main.Text("Помощь")
	main.Reply(
		main.Row(b1, b2),
		)

	return main
}

func changeCityBut(b *telebot.ReplyMarkup) telebot.Btn {
	return b.Data("Изменить город", "city")
}

func sendingSettingBut(b *telebot.ReplyMarkup) telebot.Btn {
	return b.Data("Настроить рассылку", "sending")
}

func changeCityAdmBut(user *conver.User, b *telebot.ReplyMarkup) telebot.Btn {
	return b.Data("Изменить город", "city", strconv.Itoa(user.UserId))
}

func usernameBut(i int, user conver.User, main *telebot.ReplyMarkup) telebot.Btn {
	return main.Data(user.Username, fmt.Sprintf("username%d", i), strconv.Itoa(user.UserId))
}

func namesCountBut(user *conver.User, main *telebot.ReplyMarkup) telebot.Btn {
	return main.Data(fmt.Sprintf("Имена: %dшт.", len(user.Names)), "nameSettings", strconv.Itoa(user.UserId))
}

func cancelBut(user *conver.User, main *telebot.ReplyMarkup) telebot.Btn {
	return main.Data("❌ Отмена", "cancel", strconv.Itoa(user.UserId))
}

func returnBut(user *conver.User, main *telebot.ReplyMarkup) telebot.Btn {
	return main.Data("🔙 Назад", "return", strconv.Itoa(user.UserId))
}


func addNameBut(user *conver.User, main *telebot.ReplyMarkup) telebot.Btn {
	return main.Data("✅ Добавить имя", "addName", strconv.Itoa(user.UserId))
}
