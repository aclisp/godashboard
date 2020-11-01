# vectify

`vectify` reads from stdin an HTML fragment, converts it to vecty rendering Go code, and writes to stdout.

```
$ ./vectify
<!-- Nav Item - Tables -->
      <li class="nav-item">
        <a class="nav-link" href="tables.html">
          <i class="fas fa-fw fa-table"></i>
          <span>Tables</span></a>
      </li>
^D
  // Nav Item - Tables
elem.ListItem(
  vecty.Markup(
    vecty.Class("nav-item"),
  ),
  elem.Anchor(
    vecty.Markup(
      vecty.Class("nav-link"),
      vecty.Property("href", "tables.html"),
    ),
    elem.Italic(
      vecty.Markup(
        vecty.Class("fas", "fa-fw", "fa-table"),
      ),
    ),
    elem.Span(
      vecty.Text("Tables"),
    ),
  ),
),
```

# comparison

| framework |  binding   |   event    |   two-way    |         fetch         |
| --------- | ---------- | ---------- | ------------ | --------------------- |
| vue       | `v-bind`   | `v-on`     | `v-model`    | on lifecycle or event |
| angular   | `[target]` | `(target)` | `[(target)]` | on lifecycle or event |
| react     | props      | dom events | n/a          | on lifecycle or event |
