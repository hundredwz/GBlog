function parseId(id) {
    return id.split("-")[2];
}

$(function () {
    $("button[id^='comment-reply-']").click(function () {
        let id = $(this).attr("id");
        let parentId = parseId(id);
        $("#comment-parent").val(parentId);
        $("#replyComment").modal('show');
    });

    $("#comment-new-reply").click(function () {
        $("#comment-parent").val(0);
        $("#replyComment").modal('show');
    });

    $('#comment-form').on('submit', function () {
        $(this).ajaxSubmit({
            type: 'post',
            url: '/api/comment/edit',
            success: function (result) {
                if (result.status === 1) {
                    alert("处理出错！请检查或重试！");
                } else if (result.status === 0) {
                    window.location.reload();
                } else {
                    alert("失败");
                }
            }
        });
        return false; // 阻止表单自动提交事件
    });
});