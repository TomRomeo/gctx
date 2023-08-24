package lib

import (
	"log"
	"os"
	"os/exec"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/tomromeo/gctx/pkg/structs"
	"gopkg.in/yaml.v3"
)

func ParseUserFile() []structs.User{

    var users []structs.User
    cd, err := os.UserConfigDir()
    if err != nil {
        cd = "./"
    }
    yamlFile, err := os.ReadFile(cd+"/gctx/users.yml")
    if err != nil {
        log.Printf("Could not find /gctx/users.yml: %v - %v", cd, err)
    }
    err = yaml.Unmarshal(yamlFile, &users)
    if err != nil {
        log.Printf("Could not parse yaml file: %v", err)
    }
    return users
}

func AddVimNav(list *tview.List) {
    list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        if event.Rune() == 'j' {
            list.SetCurrentItem((list.GetCurrentItem()+1)%list.GetItemCount())
            return nil
        }
        if event.Rune() == 'k' {
            list.SetCurrentItem(list.GetCurrentItem()-1)
            return nil
        }
        return event
    })

}

func ApplyUser(user structs.User) {
    exec.Command("git", "config", "--global", "--unset", "user.name").Run()
    exec.Command("git", "config", "--global", "--unset", "user.email").Run()
    exec.Command("git", "config", "--global", "--unset", "commit.gpgsign").Run()
    exec.Command("git", "config", "--global", "--unset", "user.signingkey").Run()

    exec.Command("git", "config", "--global", "user.name", user.Username).Run()
    exec.Command("git", "config", "--global", "user.email", user.Email).Run()
    if (user.GpgSign) {
        exec.Command("git", "config", "--global", "commit.gpgsign", strconv.FormatBool(user.GpgSign)).Run()
        exec.Command("git", "config", "--global", "user.signingkey", user.SigningKey).Run()
    }
}
