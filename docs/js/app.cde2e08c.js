(function(e){function t(t){for(var r,o,c=t[0],i=t[1],s=t[2],l=0,d=[];l<c.length;l++)o=c[l],Object.prototype.hasOwnProperty.call(a,o)&&a[o]&&d.push(a[o][0]),a[o]=0;for(r in i)Object.prototype.hasOwnProperty.call(i,r)&&(e[r]=i[r]);f&&f(t);while(d.length)d.shift()();return u.push.apply(u,s||[]),n()}function n(){for(var e,t=0;t<u.length;t++){for(var n=u[t],r=!0,o=1;o<n.length;o++){var c=n[o];0!==a[c]&&(r=!1)}r&&(u.splice(t--,1),e=i(i.s=n[0]))}return e}var r={},o={app:0},a={app:0},u=[];function c(e){return i.p+"js/"+({}[e]||e)+"."+{"chunk-2d0c8f97":"f33ac4be","chunk-6e0b9a1e":"eb960e8f","chunk-1f3a9a6a":"a6ff1d2d","chunk-268d8d16":"9f2ff32a"}[e]+".js"}function i(t){if(r[t])return r[t].exports;var n=r[t]={i:t,l:!1,exports:{}};return e[t].call(n.exports,n,n.exports,i),n.l=!0,n.exports}i.e=function(e){var t=[],n={"chunk-6e0b9a1e":1,"chunk-1f3a9a6a":1,"chunk-268d8d16":1};o[e]?t.push(o[e]):0!==o[e]&&n[e]&&t.push(o[e]=new Promise((function(t,n){for(var r="css/"+({}[e]||e)+"."+{"chunk-2d0c8f97":"31d6cfe0","chunk-6e0b9a1e":"ad7b79f7","chunk-1f3a9a6a":"6b61d9a7","chunk-268d8d16":"feaafe7b"}[e]+".css",a=i.p+r,u=document.getElementsByTagName("link"),c=0;c<u.length;c++){var s=u[c],l=s.getAttribute("data-href")||s.getAttribute("href");if("stylesheet"===s.rel&&(l===r||l===a))return t()}var d=document.getElementsByTagName("style");for(c=0;c<d.length;c++){s=d[c],l=s.getAttribute("data-href");if(l===r||l===a)return t()}var f=document.createElement("link");f.rel="stylesheet",f.type="text/css",f.onload=t,f.onerror=function(t){var r=t&&t.target&&t.target.src||a,u=new Error("Loading CSS chunk "+e+" failed.\n("+r+")");u.code="CSS_CHUNK_LOAD_FAILED",u.request=r,delete o[e],f.parentNode.removeChild(f),n(u)},f.href=a;var p=document.getElementsByTagName("head")[0];p.appendChild(f)})).then((function(){o[e]=0})));var r=a[e];if(0!==r)if(r)t.push(r[2]);else{var u=new Promise((function(t,n){r=a[e]=[t,n]}));t.push(r[2]=u);var s,l=document.createElement("script");l.charset="utf-8",l.timeout=120,i.nc&&l.setAttribute("nonce",i.nc),l.src=c(e);var d=new Error;s=function(t){l.onerror=l.onload=null,clearTimeout(f);var n=a[e];if(0!==n){if(n){var r=t&&("load"===t.type?"missing":t.type),o=t&&t.target&&t.target.src;d.message="Loading chunk "+e+" failed.\n("+r+": "+o+")",d.name="ChunkLoadError",d.type=r,d.request=o,n[1](d)}a[e]=void 0}};var f=setTimeout((function(){s({type:"timeout",target:l})}),12e4);l.onerror=l.onload=s,document.head.appendChild(l)}return Promise.all(t)},i.m=e,i.c=r,i.d=function(e,t,n){i.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:n})},i.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},i.t=function(e,t){if(1&t&&(e=i(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var n=Object.create(null);if(i.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var r in e)i.d(n,r,function(t){return e[t]}.bind(null,r));return n},i.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return i.d(t,"a",t),t},i.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},i.p="/CodingShare/",i.oe=function(e){throw console.error(e),e};var s=window["webpackJsonp"]=window["webpackJsonp"]||[],l=s.push.bind(s);s.push=t,s=s.slice();for(var d=0;d<s.length;d++)t(s[d]);var f=l;u.push([0,"chunk-vendors"]),n()})({0:function(e,t,n){e.exports=n("cd49")},"034f":function(e,t,n){"use strict";n("85ec")},"06eb":function(e,t,n){"use strict";n("0b8b")},"0b8b":function(e,t,n){},4313:function(e,t,n){"use strict";n("525b")},"525b":function(e,t,n){},"85ec":function(e,t,n){},cd49:function(e,t,n){"use strict";n.r(t);n("e260"),n("e6cf"),n("cca6"),n("a79d");var r=n("2b0e"),o=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{attrs:{id:"app"}},[n("div",{attrs:{id:"nav"}},[n("router-link",{attrs:{to:"/"}},[e._v("Home")]),e._v(" | "),n("router-link",{attrs:{to:"/editor"}},[e._v("Editor")]),e._v(" | "),n("router-link",{attrs:{to:"/player"}},[e._v("Player")]),e._v(" | "),n("router-link",{attrs:{to:"/list"}},[e._v("list")])],1),n("router-view")],1)},a=[],u=(n("034f"),n("2877")),c={},i=Object(u["a"])(c,o,a,!1,null,null,null),s=i.exports,l=(n("d3b7"),n("3ca3"),n("ddb0"),n("8c4f")),d=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"home-page"},[e._m(0),n("div",{staticClass:"controll-panel"},[n("router-link",{attrs:{to:"/editor"}},[n("FontAwesomeIcon",{attrs:{icon:"video"}}),e._v("録画画面へ")],1),n("router-link",{attrs:{to:"/player"}},[n("FontAwesomeIcon",{attrs:{icon:"file-video"}}),e._v("再生画面へ")],1),n("router-link",{attrs:{to:"/list"}},[n("FontAwesomeIcon",{attrs:{icon:"film"}}),e._v("動画一覧へ")],1)],1)])},f=[function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"logo-panel"},[n("h1",{staticClass:"app-logo"},[e._v("CodingShare{}")]),n("p",[e._v("本サイトはベータ版です。")]),n("p",[e._v(" 本ソフトを使用したことによるいかなる損害について、開発者は何も保証しません。 ")]),n("p",[e._v(" これに同意できない場合、あなたは本ソフトを使用することができません。 ")])])}],p=n("ecee"),h=n("c074"),v=n("ad3d");p["c"].add(h["c"],h["j"],h["d"]);var m=r["default"].extend({name:"HomePage",components:{FontAwesomeIcon:v["a"]}}),b=m,g=(n("4313"),Object(u["a"])(b,d,f,!1,null,"bbbf769a",null)),_=g.exports,y=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"not-found-page"},[n("div",[n("p",[e._v("Not found")]),n("p",[e._v("そのURLは存在しません")]),n("p",[e._v(" Top Page→ "),n("router-link",{attrs:{to:"/"}},[e._v(e._s(e.topPageUrl))])],1)])])},k=[],w=(n("99af"),r["default"].extend({name:"NotFoundPage",computed:{topPageUrl:function(){return"".concat(location.host).concat("/CodingShare/","#/")}}})),P=w,S=(n("06eb"),Object(u["a"])(P,y,k,!1,null,"45e34330",null)),C=S.exports;r["default"].use(l["a"]);var j=[{path:"/",name:"Home",component:_},{path:"/editor",name:"Editor",component:function(){return Promise.all([n.e("chunk-6e0b9a1e"),n.e("chunk-268d8d16")]).then(n.bind(null,"cdc5"))}},{path:"/player",name:"Player",component:function(){return Promise.all([n.e("chunk-6e0b9a1e"),n.e("chunk-1f3a9a6a")]).then(n.bind(null,"ac7f"))}},{path:"/list",name:"list",component:function(){return n.e("chunk-2d0c8f97").then(n.bind(null,"56bb"))}},{path:"*",name:"NotFound",component:C}],E=new l["a"]({mode:"hash",base:"/CodingShare/",routes:j}),O=E,x=n("2f62"),A=n("bfa9"),T=new A["a"]({storage:window.localStorage,key:"coding_share_vuex_key",reducer:function(e){return{speed:e.speed}},filter:function(e){return"setSpeed"===e.type}});r["default"].use(x["a"]);var N=new x["a"].Store({state:{speed:100},getters:{speed:function(e){return e.speed}},mutations:{setSpeed:function(e,t){e.speed=t}},actions:{setSpeedAction:function(e,t){var n=e.commit;n("setSpeed",t)}},modules:{},plugins:[T.plugin]}),F=n("1881"),L=n.n(F);r["default"].config.productionTip=!1,r["default"].use(L.a),new r["default"]({router:O,store:N,render:function(e){return e(s)}}).$mount("#app")}});
//# sourceMappingURL=app.cde2e08c.js.map