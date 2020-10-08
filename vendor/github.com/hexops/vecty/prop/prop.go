package prop

import "github.com/hexops/vecty"

type InputType string

const (
	TypeButton        InputType = "button"
	TypeCheckbox      InputType = "checkbox"
	TypeColor         InputType = "color"
	TypeDate          InputType = "date"
	TypeDatetime      InputType = "datetime"
	TypeDatetimeLocal InputType = "datetime-local"
	TypeEmail         InputType = "email"
	TypeFile          InputType = "file"
	TypeHidden        InputType = "hidden"
	TypeImage         InputType = "image"
	TypeMonth         InputType = "month"
	TypeNumber        InputType = "number"
	TypePassword      InputType = "password"
	TypeRadio         InputType = "radio"
	TypeRange         InputType = "range"
	TypeMin           InputType = "min"
	TypeMax           InputType = "max"
	TypeValue         InputType = "value"
	TypeStep          InputType = "step"
	TypeReset         InputType = "reset"
	TypeSearch        InputType = "search"
	TypeSubmit        InputType = "submit"
	TypeTel           InputType = "tel"
	TypeText          InputType = "text"
	TypeTime          InputType = "time"
	TypeURL           InputType = "url"
	TypeWeek          InputType = "week"
)

func Autofocus(autofocus bool) vecty.Applyer {
	return vecty.Property("autofocus", autofocus)
}

func Disabled(disabled bool) vecty.Applyer {
	return vecty.Property("disabled", disabled)
}

func Checked(checked bool) vecty.Applyer {
	return vecty.Property("checked", checked)
}

func For(id string) vecty.Applyer {
	return vecty.Property("htmlFor", id)
}

func Href(url string) vecty.Applyer {
	return vecty.Property("href", url)
}

func ID(id string) vecty.Applyer {
	return vecty.Property("id", id)
}

func Placeholder(text string) vecty.Applyer {
	return vecty.Property("placeholder", text)
}

func Src(url string) vecty.Applyer {
	return vecty.Property("src", url)
}

func Type(t InputType) vecty.Applyer {
	return vecty.Property("type", string(t))
}

func Value(v string) vecty.Applyer {
	return vecty.Property("value", v)
}

func Name(name string) vecty.Applyer {
	return vecty.Property("name", name)
}

func Alt(text string) vecty.Applyer {
	return vecty.Property("alt", text)
}
