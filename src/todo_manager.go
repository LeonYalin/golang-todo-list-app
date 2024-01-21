package src

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/leonyalin/todo-list/src/models"
	"github.com/leonyalin/todo-list/src/types"
	"github.com/leonyalin/todo-list/src/util"
)

type TodoManager struct {
	io              IOManageable
	todos           []*models.Todo
	welcome_options []types.TodoOption
}

func (tm *TodoManager) SetWelcomeOptions(options []types.TodoOption) {
	tm.welcome_options = options
}

func (tm *TodoManager) AddTodo(description string) {
	tm.todos = append(tm.todos, models.NewTodo(description))
}

func (tm *TodoManager) Print() string {
	var sb strings.Builder
	for _, t := range tm.todos {
		sb.WriteString(t.Print() + "\n")
	}
	return sb.String()
}

func (tm *TodoManager) Init() {
	tm.greetingDialog()
	tm.mainMenuDialog()
}

func (tm *TodoManager) greetingDialog() {
	tm.io.Write(fmt.Sprintf("\n%s and welcome to my %s app!\n\n", util.YellowText("Hello"), util.RedText("TODO List")))
}

func (tm *TodoManager) mainMenuDialog() {
	for i, option := range tm.welcome_options {
		tm.io.Write(util.OptionText(strconv.Itoa(i+1), option.Text) + "\n")
	}
	answer := tm.io.Question(util.ENTER_AN_OPTION_TEXT, func(ans string) bool {
		return slices.ContainsFunc(tm.welcome_options, func(item types.TodoOption) bool {
			return item.Key == ans
		})
	})
	chosenOptionIndex := slices.IndexFunc(tm.welcome_options, func(to types.TodoOption) bool {
		return to.Key == answer
	})
	tm.welcome_options[chosenOptionIndex].Handler()
}

func (tm *TodoManager) viewAllTodosDialog() {
	tm.io.Write(fmt.Sprintf("\n%s\n\n", util.YellowText("Availiable todos:")))

	if len(tm.todos) == 0 {
		tm.io.Write("No todos found.\n\n")
	} else {
		for i, todo := range tm.todos {
			tm.io.Write(fmt.Sprintf("%s\n", util.TodoText(i+1, todo)))
		}
	}

	tm.io.Write(fmt.Sprintf("%s\n", util.OptionText(types.GO_BACK, "Go back")))

	answer := tm.io.Question(util.ENTER_AN_OPTION_TEXT, func(ans string) bool {
		ansInt, err := strconv.Atoi(ans)
		return ans != "" && err == nil && ansInt >= 0 && ansInt <= len(tm.todos)
	})
	if answer == types.GO_BACK {
		tm.mainMenuDialog()
	} else {
		// TODO
	}
}

func (tm *TodoManager) addTodoDialog() {
	answer := tm.io.Question(fmt.Sprintf("%s\n", util.YellowText("Enter the todo description")), func(ans string) bool {
		return ans != ""
	})
	fmt.Println("Todo created:", models.NewTodo(answer).Print())
	tm.todos = append(tm.todos, models.NewTodo(answer))
	tm.io.Write(fmt.Sprintf("\n%s\n\n", util.GreenText("Todo successfully added!")))
	tm.mainMenuDialog()
}

func (tm *TodoManager) exitDialog() {
	tm.io.Write(fmt.Sprintf("%s\n\n", util.YellowText("Exiting.. Bye!")))
}

func NewTodoManager(io IOManageable) *TodoManager {
	tm := &TodoManager{io: io, todos: []*models.Todo{}, welcome_options: []types.TodoOption{}}
	tm.SetWelcomeOptions([]types.TodoOption{
		{Key: "1", Text: "View all todos", Handler: func(args ...any) { tm.viewAllTodosDialog() }},
		{Key: "2", Text: "Add a todo", Handler: func(args ...any) { tm.addTodoDialog() }},
		{Key: "3", Text: "Exit", Handler: func(args ...any) { tm.exitDialog() }},
	})
	return tm
}
