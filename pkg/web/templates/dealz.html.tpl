{{template "base" .}}

{{define "title"}}{{.Title}}{{end}}

{{define "content"}}
<h4>{{.Title}}</h4>

<table class="striped">
    <thead>
        <tr>
            <th>Image</th>
            <th>Name</th>
            <th>Price</th>
            <th>Savings</th>
        </tr>
    </thead>

    <tbody>
        {{range .Products}}
        <tr>
            <td><img src="https://www.ebucks.com/{{.Image}}" style="max-width: 230px; max-height: 230px"></td>
            <td><a href="{{.URL}}" target="_blank">{{.NameX}}</a></td>
            <td>{{.Price}}</td>
            <td>{{.Savings}}</td>
        </tr>
        {{else}}
        <tr>
            <td>No Deals :(</td>
        </tr>
        {{end}}
    </tbody>
</table>
{{end}}
