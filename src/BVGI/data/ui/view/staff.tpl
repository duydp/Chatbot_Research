{{define "content"}} {{$staffs := index $ "staffs"}}
<div class="jumbotron mb-0 pb-5">
    <div class="container">
        <div class="col-sm-12 col-center-block mb-5">
            <h1 align="center" class="text-primary">Staff</h1>
        </div>
        <div class="row justify-content-md-center" style="min-height: 400px;">
            <table class="table col-sm-12 table-bordered table-hover .table-striped px-0 mx-0">
                <thead class="thead-inverse col-sm-12 px-0 mx-0">
                    <tr class="row col-sm-12 px-0 mx-0">
                        <th class="col-sm-2 text-center">ID</th>
                        <th class="col-sm-4 text-center">Name</th>
                        <th class="col-sm-4 text-center">FacebookID</th>
                        <th class="col-sm-2 text-center">Action</th>
                    </tr>
                </thead>
                <tbody ic-confirm="Are you sure?" ic-target="closest tr" ic-replace-target="true">
                    {{ range $idx, $s := $staffs}}
                    <tr class="row col-sm-12 table-active px-0 mx-0">
                        <td class="col-sm-2 text-center">{{ $idx }}</td>
                        <td class="col-sm-4 text-center">{{ $s.Fullname }}&nbsp</td>
                        <td class="col-sm-4 text-center">{{ $s.FbID }}&nbsp</td>
                        <td class="col-sm-2 text-center">
                            <button class="btn btn-danger btn-lg delete" id="delete" ic-delete-from="/staff/{{$s.ID}}">Delete</button>
                        </td>
                    </tr>

                    {{end}}
                </tbody>
            </table>
        </div>
    </div>
</div>
{{end}}