/*
* Copyright (C) 2015 xiuno.com
*/

/*
参数说明：
e:
	XMLHttpRequestProgressEvent
	bubbles: false
	cancelBubble: false
	cancelable: true
	clipboardData: undefined
	currentTarget: XMLHttpRequest
	defaultPrevented: false
	eventPhase: 2
	lengthComputable: true
	loaded: 10
	path: NodeList[0]
	position: 10
	returnValue: true
	srcElement: XMLHttpRequest
	target: XMLHttpRequest
	timeStamp: 1426152501810
	total: 10
	totalSize: 10
	type: "load"
file:
	lastModified: 1422878522000
	lastModifiedDate: Mon Feb 02 2015 20:02:02 GMT+0800 (中国标准时间)
	name: "xxxxx.gif"
	size: 3001264
	type: "image/gif"
	webkitRelativePath: ""
	__proto__: File
	
// 读取本机图片


*/

if(typeof FileReader == 'undefined') console.log('FileReader undefined.');

$.fn.srcLocalFile = function(file) {
	return this.each(function() {
		var _this = this;
		/*if ($.browser.webkit) {
			this.src = window.webkitURL.createObjectURL(file); // safari Chrome8+
		} else if ($.browser.mozilla) {
			this.src = window.URL.createObjectURL(file); // FF4+
			*/
		if(window.URL) {
			this.src = window.URL.createObjectURL(file); // FF4+
		} else {
			var fr = new FileReader(); // 实例化 file reader 对象
			fr.addEventListener('load', function() { _this.src = this.result; });
			fr.readAsDataURL(file);
		}
	});
}

var FileUploader = function(fileinput, posturl, postdata) {
	this.fileinput = fileinput;
	this.posturl = posturl;
	this.postdata = postdata || {};
}

FileUploader.prototype.fileinput = null;
FileUploader.prototype.selectedfiles = null;
FileUploader.prototype.posturl = null;
FileUploader.prototype.postdata = null;
FileUploader.prototype.filename = "upfile";
FileUploader.prototype.filetype = "image/jpg"; // jpg|png|gif

FileUploader.prototype.thumb_width = 800;
//FileUploader.prototype.thumb_width2 = 100;

FileUploader.prototype.onprogress = function(file, percent) {console.log('onprogress files: %o, %d', file, percent);}
FileUploader.prototype.onselected = function(files) {console.log('selected files: %o', files); }
FileUploader.prototype.oncomplete = function(code, files) {console.log('oncomplete, code:%d, files: %o', code, files);}
FileUploader.prototype.ononce = function(file, e) {console.log('ononce: file: %o, e: %s', file, e);}
FileUploader.prototype.onerror = function(file, e) {console.log('onerror: file: %o, e: %s', file, e);}
FileUploader.prototype.onabort = function(file, e) {console.log('onabort: file: %o, e: %s', file, e);}

// 支持 HTML5 图片缩略 canvas
FileUploader.prototype.init = function(files) {
	var _this = this;
	if(!files) {
		/*
		$(_this.fileinput).off('change').on('change', function(e) {
			_this.selectedfiles = this.files;
			_this.onselected(this.files); // 调用一次
		})
		*/
		_this.file_input_change = function(e) {
			if(!_this.fileinput.value) return;
			_this.selectedfiles = this.files;
			_this.onselected(this.files); // 调用一次
		}
		_this.fileinput.removeEventListener('change', _this.file_input_change, false);
		_this.fileinput.addEventListener('change', _this.file_input_change, false);
	} else {
		_this.selectedfiles = files;
	}
	
}

FileUploader.prototype.start = function(posturl, postdata, filename) {
	if(!this.selectedfiles) return alert('Please select file!');

	if(posturl) this.posturl = posturl;
	if(postdata) this.postdata = postdata;
	if(filename) this.filename = filename;
	
	var _this = this;
	var arr = [];
	for(var i=0; i<this.selectedfiles.length; i++) {
		var file = _this.selectedfiles[i];
		+function(file) {
			// 判断文件类型，如果为图片，则缩略
			arr.push(function(callback) {
				
				var xml_http_request = function(file, jsondata) {
					//console.log("jsondata: %o",jsondata);			

					var xhr = new XMLHttpRequest();
					//xhr.timeout = 30000;
					xhr.open("POST", _this.posturl, true);
					xhr.setRequestHeader("X-Requested-With", "XMLHttpRequest");
					xhr.upload.onprogress = function(e) {
						if(e.lengthComputable) {
							var percent = Math.round(e.loaded * 100 / e.total);
							_this.onprogress(file, percent); // 回调多次
						}
					}
					// 模拟表单数据
					var formdata = new FormData();
					formdata.append(_this.filename, jsondata);
					for(k in _this.postdata) {
						formdata.append(k, _this.postdata[k]);
					}
					xhr.send(formdata);
					xhr.addEventListener('abort', function(e) {_this.onabort(file, e);});
					xhr.addEventListener('error', function(e) {_this.onerror(file, e);});
					xhr.addEventListener('load', function(e) {
						_this.ononce(file, e);
						callback(null, file); // 此处的 file 会合并到 下面的 result 数组中去
					});
					
				}
				
				var reader = new FileReader();
				reader.onload = function(e) {
					var filedata = e.target.result;
				
					// 如果是图片：
					//data:image/jpeg;base64,
					//data:application/x-msdownload;base64,
					if(_this.thumb_width > 0 && xn.substr(filedata, 0, 10) == 'data:image' && xn.substr(filedata, 0, 14) != 'data:image/gif') {
						var img = new Image();
						img.onload = function() {
							var canvas = document.createElement('canvas');
							// 等比缩放
							if(img.width > _this.thumb_width) {
								var width = _this.thumb_width;
								var height = Math.ceil((_this.thumb_width / img.width) * img.height);
							} else {
								var width = img.width;
								var height = img.height;
							}
							canvas.width = width;
							canvas.height = height;
							var ctx = canvas.getContext("2d"); 
							ctx.clearRect(0, 0, canvas.width, canvas.height); 			// canvas清屏
							ctx.drawImage(img, 0, 0, img.width, img.height, 0, 0, width, height);	// 将图像绘制到canvas上 
							var s = canvas.toDataURL(_this.filetype, 1);	
							// 如果反而变大，则直接用原来的图
							if(s.length > filedata.length) {
								s = filedata;
							}
							var data = s.substring(s.indexOf(',') + 1);
							var r = {name: file.name, width: width, height: height, data: data};
							xml_http_request(file, xn.json_encode(r));
							
						};
						img.src = filedata;
					} else {
						var s = filedata;
						var data = s.substring(s.indexOf(',') + 1);
						var r = {name: file.name, width: 0, height: 0, data: data};
						xml_http_request(file, xn.json_encode(r));
					}
				};
				reader.onerror = function(e) { console.log(e); };
				reader.readAsDataURL(file);
			});
		}(file);
	};
	async.series(arr, function(err, arr) {
		_this.fileinput.value = '';
		if(err) {
			_this.oncomplete(-1, arr);
		} else {
			_this.oncomplete(0, arr);
		}
	});
}