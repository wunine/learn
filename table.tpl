<table>
    <caption>
       <details >
            {{.Caption}}
       </details>
    </caption>
    <thead>
        {{range .THead}}
            <th>{{.}}</th>
        {{end}}
    </thead>
    <tbody>
        {{range .TBody}}
        <tr>
            <td>{{.Column1}}</td>
            <td>{{.Column2}}</td>
            <td>{{.Column3}}</td>
            <td>{{.Column4}}</td>
        </tr>
        {{end}}
    </tbody>
</table>



<style>
    table,th,td{
        border: solid;
        border-collapse: collapse;
    }
    details{
        text-align: left;
    }
</style>