// 表单快捷键提交 CTRL+ENTER   / form quick submit
$('form').keyup(function(e) {
	var jthis = $(this);
	if((e.ctrlKey && (e.which == 13 || e.which == 10)) || (e.altKey && e.which == 83)) {
		jthis.trigger('submit');
		return false;
	}
});

// 点击响应整行：方便手机浏览  / check response line
$('.tap').on('click', function(e) {
	var href = $(this).attr('href') || $(this).data('href');
	if(e.target.nodeName == 'INPUT') return true;
	if($(window).width() > 992) return;
	if(e.ctrlKey) {
		window.open(href);
		return false;
	} else {
		window.location = href;
	}
});
// 点击响应整行：导航栏下拉菜单   / check response line
$('ul.nav > li').on('click', function(e) {
	var jthis = $(this);
	var href = jthis.children('a').attr('href');
	if(e.ctrlKey) {
		window.open(href);
		return false;
	}
});
// 点击响应整行：，但是不响应 checkbox 的点击  / check response line, without checkbox
$('.thread input[type="checkbox"]').parents('td').on('click', function(e) {
	e.stopPropagation();
})

// 版主管理：删除 / moderator : delete
/*
$('.mod-button button.delete').on('click', function() {
	var modtid = $('input[name="modtid"]').checked();
	if(modtid.length == 0) return $.alert(lang.please_choose_thread);
	$.confirm(xn.lang('confirm_delete_thread', {n:modtid.length}), function() {
		var tids = xn.implode('_', modtid);
		$.xpost(xn.url('mod-delete-'+tids), function(code, message) {
			if(code != 0) return $.alert(message);
			$.alert(message).delay(1000).location('');
		});
	});
})
*/

// 版主管理：移动 / moderator : move
/*
$('.mod-button button.move').on('click', function() {
	var modtid = $('input[name="modtid"]').checked();
	if(modtid.length == 0) return $.alert(lang.please_choose_thread);
	var select = xn.form_select('fid', forumarr, fid);
	$.confirm(lang.move_forum, function() {
		var tids = xn.implode('_', modtid);
		var newfid = $('select[name="fid"]').val();
		$.xpost(xn.url('mod-move-'+tids+'-'+newfid), function(code, message) {
			if(code != 0) return $.alert(message);
			$.alert(message).delay(1000).location('');
		});
	}, {'body': '<p>'+lang.choose_move_forum+'：'+select+'</p>'});
})
*/

// 版主管理：置顶
/*
$('.mod-button button.top').on('click', function() {
	var modtid = $('input[name="modtid"]').checked();
	if(modtid.length == 0) return $.alert(lang.please_choose_thread);
	var lang_top = {"0": lang.top_0, "1": lang.top_1};
	if(gid == 1) lang_top["3"] = lang.top_3; //  || gid == 2
	var radios = xn.form_radio('top', lang_top);
	$.confirm(lang.top_thread, function() {
		var tids = xn.implode('_', modtid);
		var top = $('input[name="top"]').checked();
		var postdata = {top: top};
		$.xpost(xn.url('mod-top-'+tids), postdata, function(code, message) {
			if(code != 0) return $.alert(message);
			$.alert(message).delay(1000).location('');
		});
	}, {'body': '<p>'+lang.top_range+'：'+radios+'</p>'});
})
*/

// 版主管理：关闭/开启
/*
$('.mod-button button._close').on('click', function() {
	var modtid = $('input[name="modtid"]').checked();
	if(modtid.length == 0) return $.alert(lang.please_choose_thread);
	var radios = xn.form_radio('close', {"0": lang.open, "1": lang.close});
	$.confirm(lang.close_thread, function() {
		var tids = xn.implode('_', modtid);
		var close = $('input[name="close"]').checked();
		var postdata = {close: close};
		$.xpost(xn.url('mod-close-'+tids), postdata, function(code, message) {
			if(code != 0) return $.alert(message);
			$.alert(message).delay(1000).location('');
		});
	}, {'body': '<p>'+lang.close_status+'：'+radios+'</p>'});
})
*/

// 确定框 / confirm / GET / POST
// <a href="1.php" data-confirm-text="确定删除？" class="confirm">删除</a>
// <a href="1.php" data-method="post" data-confirm-text="确定删除？" class="confirm">删除</a>
$('a.confirm').on('click', function() {
	var jthis = $(this);
	var text = jthis.data('confirm-text');
	$.confirm(text, function() {
		var method = xn.strtolower(jthis.data('method'));
		var href = jthis.data('href') || jthis.attr('href');
		if(method == 'post') {
			$.xpost(href, function(code, message) {
				if(code == 0) {
					window.location.reload();
				} else {
					alert(message);					
				}
			});
		} else {
			//window.location = jthis.attr('href');
		}
	})
	return false;
});

// 选中所有 / check all
// <input class="checkall" data-target=".tid" />
$('input.checkall').on('click', function() {
	var jthis = $(this);
	var target = jthis.data('target');
	jtarget = $(target);
	jtarget.prop('checked', this.checked);
});

/*
jmobile_collapsing_bavbar = $('#mobile_collapsing_bavbar');
jmobile_collapsing_bavbar.on('touchstart', function(e) {
	//var h = $(window).height() - 120;
	var h = 350;
	jmobile_collapsing_bavbar.css('overflow-y', 'auto').css('max-height', h+'px');
	e.stopPropagation();
});
jmobile_collapsing_bavbar.on('touchmove', function(e) {
	//e.stopPropagation();
	//e.stopImmediatePropagation();
});*/

// hack: history.back() cannot back, go to the index
//$('.xn-back').on('click', function() {
	//$('.xn-back').delay(10000).location('./');
	//return false;
//});



// 删除帖子 / Delete post
$('body').on('click', '.post_delete', function() {
	var jthis = $(this);
	var href = jthis.data('href');
	var isfirst = jthis.attr('isfirst');
	if(window.confirm(lang.confirm_delete)) {
		$.xpost(href, function(code, message) {
			var isfirst = jthis.attr('isfirst');
			if(code == 0) {
				if(isfirst == '1') {
					$.location('<?php echo url("forum-$fid");?>');
				} else {
					// 删掉楼层
					jthis.parents('.post').remove();
					// 回复数 -1
					var jposts = $('.posts');
					jposts.html(xn.intval(jposts.html()) - 1);
				}
			} else {
				$.alert(message);
			}
		});
	}
	return false;
});

// 引用 / Quote
$('body').on('click', '.post_reply', function() {
	var jthis = $(this);
	var tid = jthis.data('tid');
	var pid = jthis.data('pid');
	var jmessage = $('#message');
	var jli = jthis.closest('.post');
	var jpostlist = jli.closest('.postlist');
	var jadvanced_reply = $('#advanced_reply');
	var jform = $('#quick_reply_form');
	if(jli.hasClass('quote')) {
		jli.removeClass('quote');
		jform.find('input[name="quotepid"]').val(0);
		jadvanced_reply.attr('href', xn.url('post-create-'+tid));
	} else {
		jpostlist.find('.post').removeClass('quote');
		jli.addClass('quote');
		var s = jmessage.val();
		jform.find('input[name="quotepid"]').val(pid);
		jadvanced_reply.attr('href', xn.url('post-create-'+tid+'-0-'+pid));
	}
	jmessage.focus();
	return false;
});