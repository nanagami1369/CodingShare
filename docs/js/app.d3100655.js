(function(e){function t(t){for(var r,o,u=t[0],i=t[1],l=t[2],s=0,d=[];s<u.length;s++)o=u[s],Object.prototype.hasOwnProperty.call(a,o)&&a[o]&&d.push(a[o][0]),a[o]=0;for(r in i)Object.prototype.hasOwnProperty.call(i,r)&&(e[r]=i[r]);f&&f(t);while(d.length)d.shift()();return c.push.apply(c,l||[]),n()}function n(){for(var e,t=0;t<c.length;t++){for(var n=c[t],r=!0,o=1;o<n.length;o++){var u=n[o];0!==a[u]&&(r=!1)}r&&(c.splice(t--,1),e=i(i.s=n[0]))}return e}var r={},o={app:0},a={app:0},c=[];function u(e){return i.p+"js/"+({}[e]||e)+"."+{"chunk-2d0c8f97":"0c14bc47","chunk-7129d4a2":"378af26e","chunk-142b87dc":"111b88eb","chunk-220214c6":"9682474a"}[e]+".js"}function i(t){if(r[t])return r[t].exports;var n=r[t]={i:t,l:!1,exports:{}};return e[t].call(n.exports,n,n.exports,i),n.l=!0,n.exports}i.e=function(e){var t=[],n={"chunk-7129d4a2":1,"chunk-142b87dc":1,"chunk-220214c6":1};o[e]?t.push(o[e]):0!==o[e]&&n[e]&&t.push(o[e]=new Promise((function(t,n){for(var r="css/"+({}[e]||e)+"."+{"chunk-2d0c8f97":"31d6cfe0","chunk-7129d4a2":"64ca2aae","chunk-142b87dc":"03f6d52a","chunk-220214c6":"1c41c1bd"}[e]+".css",a=i.p+r,c=document.getElementsByTagName("link"),u=0;u<c.length;u++){var l=c[u],s=l.getAttribute("data-href")||l.getAttribute("href");if("stylesheet"===l.rel&&(s===r||s===a))return t()}var d=document.getElementsByTagName("style");for(u=0;u<d.length;u++){l=d[u],s=l.getAttribute("data-href");if(s===r||s===a)return t()}var f=document.createElement("link");f.rel="stylesheet",f.type="text/css",f.onload=t,f.onerror=function(t){var r=t&&t.target&&t.target.src||a,c=new Error("Loading CSS chunk "+e+" failed.\n("+r+")");c.code="CSS_CHUNK_LOAD_FAILED",c.request=r,delete o[e],f.parentNode.removeChild(f),n(c)},f.href=a;var p=document.getElementsByTagName("head")[0];p.appendChild(f)})).then((function(){o[e]=0})));var r=a[e];if(0!==r)if(r)t.push(r[2]);else{var c=new Promise((function(t,n){r=a[e]=[t,n]}));t.push(r[2]=c);var l,s=document.createElement("script");s.charset="utf-8",s.timeout=120,i.nc&&s.setAttribute("nonce",i.nc),s.src=u(e);var d=new Error;l=function(t){s.onerror=s.onload=null,clearTimeout(f);var n=a[e];if(0!==n){if(n){var r=t&&("load"===t.type?"missing":t.type),o=t&&t.target&&t.target.src;d.message="Loading chunk "+e+" failed.\n("+r+": "+o+")",d.name="ChunkLoadError",d.type=r,d.request=o,n[1](d)}a[e]=void 0}};var f=setTimeout((function(){l({type:"timeout",target:s})}),12e4);s.onerror=s.onload=l,document.head.appendChild(s)}return Promise.all(t)},i.m=e,i.c=r,i.d=function(e,t,n){i.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:n})},i.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},i.t=function(e,t){if(1&t&&(e=i(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var n=Object.create(null);if(i.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var r in e)i.d(n,r,function(t){return e[t]}.bind(null,r));return n},i.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return i.d(t,"a",t),t},i.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},i.p="/CodingShare/",i.oe=function(e){throw console.error(e),e};var l=window["webpackJsonp"]=window["webpackJsonp"]||[],s=l.push.bind(l);l.push=t,l=l.slice();for(var d=0;d<l.length;d++)t(l[d]);var f=s;c.push([0,"chunk-vendors"]),n()})({0:function(e,t,n){e.exports=n("cd49")},"034f":function(e,t,n){"use strict";n("85ec")},"06eb":function(e,t,n){"use strict";n("0b8b")},"0b8b":function(e,t,n){},4313:function(e,t,n){"use strict";n("525b")},"525b":function(e,t,n){},"85ec":function(e,t,n){},cd49:function(e,t,n){"use strict";n.r(t);n("e260"),n("e6cf"),n("cca6"),n("a79d");var r=n("2b0e"),o=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{attrs:{id:"app"}},[n("div",{attrs:{id:"nav"}},[n("router-link",{attrs:{to:"/"}},[e._v("Home")]),e._v(" | "),n("router-link",{attrs:{to:"/editor"}},[e._v("Editor")]),e._v(" | "),n("router-link",{attrs:{to:"/player"}},[e._v("Player")]),e._v(" | "),n("router-link",{attrs:{to:"/list"}},[e._v("list")])],1),n("router-view")],1)},a=[],c=(n("034f"),n("2877")),u={},i=Object(c["a"])(u,o,a,!1,null,null,null),l=i.exports,s=(n("d3b7"),n("3ca3"),n("ddb0"),n("8c4f")),d=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"home-page"},[e._m(0),n("div",{staticClass:"controll-panel"},[n("router-link",{attrs:{to:"/editor"}},[n("FontAwesomeIcon",{attrs:{icon:"video"}}),e._v("録画画面へ")],1),n("router-link",{attrs:{to:"/player"}},[n("FontAwesomeIcon",{attrs:{icon:"file-video"}}),e._v("再生画面へ")],1),n("router-link",{attrs:{to:"/list"}},[n("FontAwesomeIcon",{attrs:{icon:"film"}}),e._v("動画一覧へ")],1)],1)])},f=[function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"logo-panel"},[n("h1",{staticClass:"app-logo"},[e._v("CodingShare{}")]),n("p",[e._v("本サイトはベータ版です。")]),n("p",[e._v(" 本ソフトを使用したことによるいかなる損害について、開発者は何も保証しません。 ")]),n("p",[e._v(" これに同意できない場合、あなたは本ソフトを使用することができません。 ")])])}],p=n("ecee"),h=n("c074"),v=n("ad3d");p["c"].add(h["c"],h["j"],h["d"]);var m=r["default"].extend({name:"HomePage",components:{FontAwesomeIcon:v["a"]}}),b=m,g=(n("4313"),Object(c["a"])(b,d,f,!1,null,"bbbf769a",null)),_=g.exports,y=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"not-found-page"},[n("div",[n("p",[e._v("Not found")]),n("p",[e._v("そのURLは存在しません")]),n("p",[e._v(" Top Page→ "),n("router-link",{attrs:{to:"/"}},[e._v(e._s(e.topPageUrl))])],1)])])},k=[],w=(n("99af"),r["default"].extend({name:"NotFoundPage",computed:{topPageUrl:function(){return"".concat(location.host).concat("/CodingShare/","#/")}}})),P=w,C=(n("06eb"),Object(c["a"])(P,y,k,!1,null,"45e34330",null)),j=C.exports;r["default"].use(s["a"]);var E=[{path:"/",name:"Home",component:_},{path:"/editor",name:"Editor",component:function(){return Promise.all([n.e("chunk-7129d4a2"),n.e("chunk-220214c6")]).then(n.bind(null,"cdc5"))}},{path:"/player",name:"Player",component:function(){return Promise.all([n.e("chunk-7129d4a2"),n.e("chunk-142b87dc")]).then(n.bind(null,"ac7f"))}},{path:"/list",name:"list",component:function(){return n.e("chunk-2d0c8f97").then(n.bind(null,"56bb"))}},{path:"*",name:"NotFound",component:j}],O=new s["a"]({mode:"hash",base:"/CodingShare/",routes:E}),S=O,x=n("2f62");r["default"].use(x["a"]);var A=new x["a"].Store({state:{},mutations:{},actions:{},modules:{}}),T=n("1881"),N=n.n(T);r["default"].config.productionTip=!1,r["default"].use(N.a),new r["default"]({router:S,store:A,render:function(e){return e(l)}}).$mount("#app")}});
//# sourceMappingURL=app.d3100655.js.map