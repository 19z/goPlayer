{{template "header.tpl"}}
<div class="container">
    <div class="row">
        <div class="col-md-8">
            <div class="panel panel-default">
                <div class="panel-heading">共{{.Number}}条记录</div>
                <div class="panel-body">
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
                </div>
            </div>


        </div>
        <div class="col-md-4">
            <div class="panel panel-default">
                <div class="panel-heading">新增</div>
                <div class="panel-body">
                    <form method="post">
                        <div class="form-group">
                            <label>名称</label>
                            <input name="name" type="text" class="form-control">
                        </div>
                        <div class="form-group">
                            <label>类型</label>
                            <select name="type" class="form-control">
                                <option value="local">本地</option>
                            </select>
                        </div>
                        <div class="form-group">
                            <label>地址</label>
                            <input name="path" type="text" class="form-control">
                        </div>
                        <input type="submit" class="btn btn-default" value="保存">
                    </form>
                </div>
            </div>
        </div>
    </div>

</div>
{{template "footer.tpl"}}