$(function () {
    $('#article-edit').on('submit', function () {
        $(this).ajaxSubmit({
            type: 'post',
            url: '/api/article/edit',
            success: function (result) {
                if (result.status === 1) {
                    alert("处理出错！请检查或重试！");
                } else if (result.status === 0) {
                    alert("成功");
                    window.location.href = "/admin/article/list"
                } else {
                    alert("更新失败");
                }
            }
        });
        return false; // 阻止表单自动提交事件
    });

    $("button[id^='article-delete-']").click(function () {
        let id = $(this).attr("id");
        let cid = parseId(id);
        let msg = confirm("确定要删除本文章？");
        if (msg === true) {
            delArticle(cid);
        }
    });

    $("button[id^='article-publish-']").click(function () {
        let id = $(this).attr("id");
        let cid = parseId(id);
        let msg = confirm("确定要发表本文章？");
        if (msg === true) {
            changeArticleStatus(cid, "publish");
        }
    });

    $("button[id^='article-draft-']").click(function () {
        let id = $(this).attr("id");
        let cid = parseId(id);
        let msg = confirm("确定要改为草稿状态？");
        if (msg === true) {
            changeArticleStatus(cid, "draft");
        }
    });

    $("button[id^='comment-delete-']").click(function () {
        let id = $(this).attr("id");
        let coid = parseId(id);
        let msg = confirm("确定要删除本评论？");
        if (msg === true) {
            delComment(coid);
        }
    });

    $("button[id^='comment-audit-']").click(function () {
        let id = $(this).attr("id");
        let coid = parseId(id);
        let msg = confirm("确定修改为审核状态?");
        if (msg === true) {
            changeCommentStatus(coid, "waiting");
        }
    });

    $("button[id^='comment-approve-']").click(function () {
        let id = $(this).attr("id");
        let coid = parseId(id);
        let msg = confirm("确定修改为通过状态?");
        if (msg === true) {
            changeCommentStatus(coid, "approved");
        }
    });

    $("button[id^='comment-reply-']").click(function () {
        let id = $(this).attr("id");
        let parentId = parseId(id);
        $("#comment-parent").val(parentId);
        $("#comment-cid").val($(this).val());
        $("#replyComment").modal('show');
    });

    $("button[id^='comment-edit-']").click(function () {
        let id = $(this).attr("id");
        let coid = parseId(id);
        $("#comment-coid").val(coid);
        getComment(coid);
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
                    alert("成功");
                    window.location.reload();
                } else {
                    alert("失败");
                }
            }
        });
        return false; // 阻止表单自动提交事件
    });

    $('#page-edit').on('submit', function () {
        $(this).ajaxSubmit({
            type: 'post',
            url: '/api/page/edit',
            success: function (result) {
                if (result.status === 1) {
                    alert("处理出错！请检查或重试！");
                } else if (result.status === 0) {
                    alert("成功");
                    window.location.href = "/admin/page/list"
                } else {
                    alert("更新失败");
                }
            }
        });
        return false; // 阻止表单自动提交事件
    });

    $("button[id^='page-delete-']").click(function () {
        let id = $(this).attr("id");
        let cid = parseId(id);
        let msg = confirm("确定要删除本页面？");
        if (msg === true) {
            delPage(cid);
        }
    });

    $("button[id^='page-publish-']").click(function () {
        let id = $(this).attr("id");
        let cid = parseId(id);
        let msg = confirm("确定要发表本页面？");
        if (msg === true) {
            changePageStatus(cid, "publish");
        }
    });

    $("button[id^='page-draft-']").click(function () {
        let id = $(this).attr("id");
        let cid = parseId(id);
        let msg = confirm("确定要改为草稿状态？");
        if (msg === true) {
            changePageStatus(cid, "draft");
        }
    });

    $('#tag-edit').on('submit', function () {
        $(this).ajaxSubmit({
            type: 'post',
            url: '/api/tag/edit',
            success: function (result) {
                if (result.status === 1) {
                    alert("处理出错！请检查或重试！");
                } else if (result.status === 0) {
                    alert("成功");
                    window.location.href = "/admin/tag/list"
                } else {
                    alert("更新失败");
                }
            }
        });
        return false; // 阻止表单自动提交事件
    });

    $('#category-edit').on('submit', function () {
        $(this).ajaxSubmit({
            type: 'post',
            url: '/api/category/edit',
            success: function (result) {
                if (result.status === 1) {
                    alert("处理出错！请检查或重试！");
                } else if (result.status === 0) {
                    alert("成功");
                    window.location.href = "/admin/category/list"
                } else {
                    alert("更新失败");
                }
            }
        });
        return false; // 阻止表单自动提交事件
    });

    $("button[id^='category-delete-']").click(function () {
        let id = $(this).attr("id");
        let cid = parseId(id);
        let msg = confirm("确定要删除本分类？")
        if (msg === true) {
            delCategory(cid);
        }
    });

    $("button[id^='tag-delete']").click(function () {
        let mid = $('input[id=mid]').val();
        let msg = confirm("确定要删除本信息？")
        if (msg === true) {
            delTag(mid);
        }
        return false;
    });

    $('#blog-config').on('submit', function () {
        $(this).ajaxSubmit({
            type: 'post',
            url: '/api/setting/blog',
            success: function (result) {
                if (result.status === 1) {
                    alert("处理出错！请检查或重试！");
                } else if (result.status === 0) {
                    alert("成功");
                    window.location.href = "/admin/setting/blog"
                } else {
                    alert("更新失败");
                }
            }
        });
        return false; // 阻止表单自动提交事件
    });

    $('#user-config').on('submit', function () {
        $(this).ajaxSubmit({
            type: 'post',
            url: '/api/setting/user',
            success: function (result) {
                if (result.status === 1) {
                    alert("处理出错！请检查或重试！");
                } else if (result.status === 0) {
                    alert("成功");
                    window.location.href = "/admin/setting/user"
                } else {
                    alert("更新失败");
                }
            }
        });
        return false; // 阻止表单自动提交事件
    });

});

function parseId(id) {
    return id.split("-")[2];
}

function delArticle(id) {
    $.ajax({
        type: "get",
        url: "/api/article/delete",
        data: {
            'cid': id
        },
        success: function (result) {
            if (result.status === 1) {
                alert("处理出错！请检查或重试！")
            } else if (result.status === 0) {
                alert("删除成功");
                window.location.reload();
            } else {
                alert("系统内部错误");
            }
        }
    });
}

function changeArticleStatus(id, status) {
    $.ajax({
        type: "post",
        url: "/api/article/status",
        data: {
            'cid': id,
            "status": status
        },
        success: function (result) {
            if (result.status === 1) {
                alert("处理出错！请检查或重试！")
            } else if (result.status === 0) {
                alert("修改成功");
                window.location.reload();
            } else {
                alert("系统内部错误");
            }
        }
    });
}

function changePageStatus(id, status) {
    $.ajax({
        type: "post",
        url: "/api/page/status",
        data: {
            'cid': id,
            "status": status
        },
        success: function (result) {
            if (result.status === 1) {
                alert("处理出错！请检查或重试！")
            } else if (result.status === 0) {
                alert("修改成功");
                window.location.reload();
            } else {
                alert("系统内部错误");
            }
        }
    });
}

function delPage(id) {
    $.ajax({
        type: "get",
        url: "/api/page/delete",
        data: {
            'cid': id
        },
        success: function (result) {
            if (result.status === 1) {
                alert("处理出错！请检查或重试！")
            } else if (result.status === 0) {
                alert("删除成功");
                window.location.reload();
            } else {
                alert("系统内部错误");
            }
        }
    });
}

function delCategory(id) {
    $.ajax({
        type: "get",
        url: "/api/category/delete",
        data: {
            'mid': id
        },
        success: function (result) {
            if (result.status === 1) {
                alert("处理出错！请检查或重试！")
            } else if (result.status === 0) {
                alert("删除成功");
                window.location.reload();
            } else {
                alert("系统内部错误");
            }
        }
    });
}

function delTag(id) {
    $.ajax({
        type: "get",
        url: "/api/tag/delete",
        data: {
            'mid': id
        },
        success: function (result) {
            if (result.status === 1) {
                alert("删除失败")
            } else if (result.status === 0) {
                alert("删除成功");
                window.location.reload();
            } else {
                alert("系统内部错误");
            }
        }
    });
}

function delComment(id) {
    $.ajax({
        type: "get",
        url: "/api/comment/delete",
        data: {
            'coid': id
        },
        success: function (result) {
            if (result.status === 1) {
                alert("处理出错！请检查或重试！")
            } else if (result.status === 0) {
                alert("删除成功");
                window.location.reload();
            } else {
                alert("系统内部错误");
            }
        }
    });
}

function changeCommentStatus(id, status) {
    $.ajax({
        type: "post",
        url: "/api/comment/status",
        data: {
            'coid': id,
            "status": status
        },
        success: function (result) {
            if (result.status === 1) {
                alert("处理出错！请检查或重试！")
            } else if (result.status === 0) {
                alert("修改成功");
                window.location.reload();
            } else {
                alert("系统内部错误");
            }
        }
    });
}

function updateTag(slug) {
    $('#tagEdit').html("编辑标签");
    $("#tag-delete").css('display', 'block');
    $.ajax({
        url: "/api/tag?slug=" + slug,
        async: true,
        success: function (result) {
            if (result.status === 1) {
                alert("处理出错！请检查或重试！");
            } else if (result.status === 0) {
                $("#mid").val(result.payload.mid);
                $("#tagName").val(result.payload.name);
                $("#tagSlug").val(result.payload.slug);
                $("#tagDesc").val(result.payload.description);
            } else {
                alert("获取失败");
            }
        }
    });

}

function getComment(coid) {
    $.ajax({
        url: "/api/comment?coid=" + coid,
        async: true,
        success: function (result) {
            if (result.status === 1) {
                alert("处理出错！请检查或重试！");
            } else if (result.status === 0) {
                $("#comment-username").val(result.payload.author);
                $("#comment-mail").val(result.payload.mail);
                $("#comment-url").val(result.payload.url);
                $("#comment-text").val(result.payload.text);
                $("#comment-created").val(result.payload.created);
            } else {
                alert("获取失败");
            }
        }
    });
}