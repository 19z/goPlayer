{{template "header.tpl"}}
共{{.Number}}条记录
<table class="table">
    <thead>
    <tr>
        <th>ID</th>
        <th>名称</th>
        <th>创建时间</th>
    </tr>
    </thead>
    <tbody>
    {{range $i,$driver:=.Drivers}}
        <tr>
            <td>{{$driver.Id}}</td>
            <td><a href="/files/{{$driver.Id}}/~/">{{$driver.Name}}</a></td>
            <td>{{$driver.CreatedTime}}</td>
        </tr>
    {{end}}
    </tbody>
</table>
{{template "footer.tpl"}}