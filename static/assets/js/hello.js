$(function() {
    $("button[type='submit']").click(function() {
        $.post("/api/hello", {
            name: $("input[name='name']").val(),
            age: $("input[name='age']").val()
        }, function(data) {
            $("table.hello-table").append(
                "<tr>" +
                    "<td>" + data.name + "</td>" +
                    "<td>" + data.age + "</td>" +
                "</tr>"
            )
        });
        return false;
    });
});
