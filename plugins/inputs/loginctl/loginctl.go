package loginctl

import (
	lctl "github.com/corporateanon/loginctl"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

type Loginctl struct {
	loginctl lctl.ILoginctl
}

func (_ *Loginctl) Description() string {
	return "Read desktop user activity"
}

func (_ *Loginctl) SampleConfig() string { return "" }

func (m *Loginctl) Init() error {
	return nil
}

func (s *Loginctl) Gather(acc telegraf.Accumulator) error {

	sessionInfo, err := s.loginctl.GetSessionInfo()
	if err != nil {
		return err
	}

	for user, active := range sessionInfo.UserActivities {
		var activity int64
		if active {
			activity = 1
		} else {
			activity = 0
		}
		tags := map[string]string{
			"user": user,
		}
		fields := map[string]interface{}{
			"active": activity,
		}
		acc.AddFields("loginctl", fields, tags)
	}

	return nil
}

func init() {
	inputs.Add("loginctl", func() telegraf.Input {
		loginctl, err := lctl.NewFromRegularUsers()
		if err != nil {
			panic(err)
		}
		return &Loginctl{
			loginctl: loginctl,
		}
	})
}
