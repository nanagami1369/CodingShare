(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-64f9a27d"],{"0d65":function(t,e,n){"use strict";n("48ad")},"10f9":function(t,e,n){},"133e":function(t,e,n){},"3f30":function(t,e,n){"use strict";n("10f9")},"48ad":function(t,e,n){},"4f60":function(t,e,n){"use strict";n("133e")},5899:function(t,e){t.exports="\t\n\v\f\r                　\u2028\u2029\ufeff"},"58a8":function(t,e,n){var r=n("1d80"),a=n("5899"),i="["+a+"]",o=RegExp("^"+i+i+"*"),u=RegExp(i+i+"*$"),s=function(t){return function(e){var n=String(r(e));return 1&t&&(n=n.replace(o,"")),2&t&&(n=n.replace(u,"")),n}};t.exports={start:s(1),end:s(2),trim:s(3)}},7156:function(t,e,n){var r=n("861d"),a=n("d2bb");t.exports=function(t,e,n){var i,o;return a&&"function"==typeof(i=e.constructor)&&i!==n&&r(o=i.prototype)&&o!==n.prototype&&a(t,o),t}},"96cf":function(t,e,n){var r=function(t){"use strict";var e,n=Object.prototype,r=n.hasOwnProperty,a="function"===typeof Symbol?Symbol:{},i=a.iterator||"@@iterator",o=a.asyncIterator||"@@asyncIterator",u=a.toStringTag||"@@toStringTag";function s(t,e,n){return Object.defineProperty(t,e,{value:n,enumerable:!0,configurable:!0,writable:!0}),t[e]}try{s({},"")}catch(U){s=function(t,e,n){return t[e]=n}}function c(t,e,n,r){var a=e&&e.prototype instanceof v?e:v,i=Object.create(a.prototype),o=new P(r||[]);return i._invoke=S(t,n,o),i}function d(t,e,n){try{return{type:"normal",arg:t.call(e,n)}}catch(U){return{type:"throw",arg:U}}}t.wrap=c;var f="suspendedStart",l="suspendedYield",h="executing",m="completed",g={};function v(){}function w(){}function p(){}var y={};y[i]=function(){return this};var b=Object.getPrototypeOf,T=b&&b(b(N([])));T&&T!==n&&r.call(T,i)&&(y=T);var x=p.prototype=v.prototype=Object.create(y);function C(t){["next","throw","return"].forEach((function(e){s(t,e,(function(t){return this._invoke(e,t)}))}))}function M(t,e){function n(a,i,o,u){var s=d(t[a],t,i);if("throw"!==s.type){var c=s.arg,f=c.value;return f&&"object"===typeof f&&r.call(f,"__await")?e.resolve(f.__await).then((function(t){n("next",t,o,u)}),(function(t){n("throw",t,o,u)})):e.resolve(f).then((function(t){c.value=t,o(c)}),(function(t){return n("throw",t,o,u)}))}u(s.arg)}var a;function i(t,r){function i(){return new e((function(e,a){n(t,r,e,a)}))}return a=a?a.then(i,i):i()}this._invoke=i}function S(t,e,n){var r=f;return function(a,i){if(r===h)throw new Error("Generator is already running");if(r===m){if("throw"===a)throw i;return D()}n.method=a,n.arg=i;while(1){var o=n.delegate;if(o){var u=E(o,n);if(u){if(u===g)continue;return u}}if("next"===n.method)n.sent=n._sent=n.arg;else if("throw"===n.method){if(r===f)throw r=m,n.arg;n.dispatchException(n.arg)}else"return"===n.method&&n.abrupt("return",n.arg);r=h;var s=d(t,e,n);if("normal"===s.type){if(r=n.done?m:l,s.arg===g)continue;return{value:s.arg,done:n.done}}"throw"===s.type&&(r=m,n.method="throw",n.arg=s.arg)}}}function E(t,n){var r=t.iterator[n.method];if(r===e){if(n.delegate=null,"throw"===n.method){if(t.iterator["return"]&&(n.method="return",n.arg=e,E(t,n),"throw"===n.method))return g;n.method="throw",n.arg=new TypeError("The iterator does not provide a 'throw' method")}return g}var a=d(r,t.iterator,n.arg);if("throw"===a.type)return n.method="throw",n.arg=a.arg,n.delegate=null,g;var i=a.arg;return i?i.done?(n[t.resultName]=i.value,n.next=t.nextLoc,"return"!==n.method&&(n.method="next",n.arg=e),n.delegate=null,g):i:(n.method="throw",n.arg=new TypeError("iterator result is not an object"),n.delegate=null,g)}function k(t){var e={tryLoc:t[0]};1 in t&&(e.catchLoc=t[1]),2 in t&&(e.finallyLoc=t[2],e.afterLoc=t[3]),this.tryEntries.push(e)}function _(t){var e=t.completion||{};e.type="normal",delete e.arg,t.completion=e}function P(t){this.tryEntries=[{tryLoc:"root"}],t.forEach(k,this),this.reset(!0)}function N(t){if(t){var n=t[i];if(n)return n.call(t);if("function"===typeof t.next)return t;if(!isNaN(t.length)){var a=-1,o=function n(){while(++a<t.length)if(r.call(t,a))return n.value=t[a],n.done=!1,n;return n.value=e,n.done=!0,n};return o.next=o}}return{next:D}}function D(){return{value:e,done:!0}}return w.prototype=x.constructor=p,p.constructor=w,w.displayName=s(p,u,"GeneratorFunction"),t.isGeneratorFunction=function(t){var e="function"===typeof t&&t.constructor;return!!e&&(e===w||"GeneratorFunction"===(e.displayName||e.name))},t.mark=function(t){return Object.setPrototypeOf?Object.setPrototypeOf(t,p):(t.__proto__=p,s(t,u,"GeneratorFunction")),t.prototype=Object.create(x),t},t.awrap=function(t){return{__await:t}},C(M.prototype),M.prototype[o]=function(){return this},t.AsyncIterator=M,t.async=function(e,n,r,a,i){void 0===i&&(i=Promise);var o=new M(c(e,n,r,a),i);return t.isGeneratorFunction(n)?o:o.next().then((function(t){return t.done?t.value:o.next()}))},C(x),s(x,u,"Generator"),x[i]=function(){return this},x.toString=function(){return"[object Generator]"},t.keys=function(t){var e=[];for(var n in t)e.push(n);return e.reverse(),function n(){while(e.length){var r=e.pop();if(r in t)return n.value=r,n.done=!1,n}return n.done=!0,n}},t.values=N,P.prototype={constructor:P,reset:function(t){if(this.prev=0,this.next=0,this.sent=this._sent=e,this.done=!1,this.delegate=null,this.method="next",this.arg=e,this.tryEntries.forEach(_),!t)for(var n in this)"t"===n.charAt(0)&&r.call(this,n)&&!isNaN(+n.slice(1))&&(this[n]=e)},stop:function(){this.done=!0;var t=this.tryEntries[0],e=t.completion;if("throw"===e.type)throw e.arg;return this.rval},dispatchException:function(t){if(this.done)throw t;var n=this;function a(r,a){return u.type="throw",u.arg=t,n.next=r,a&&(n.method="next",n.arg=e),!!a}for(var i=this.tryEntries.length-1;i>=0;--i){var o=this.tryEntries[i],u=o.completion;if("root"===o.tryLoc)return a("end");if(o.tryLoc<=this.prev){var s=r.call(o,"catchLoc"),c=r.call(o,"finallyLoc");if(s&&c){if(this.prev<o.catchLoc)return a(o.catchLoc,!0);if(this.prev<o.finallyLoc)return a(o.finallyLoc)}else if(s){if(this.prev<o.catchLoc)return a(o.catchLoc,!0)}else{if(!c)throw new Error("try statement without catch or finally");if(this.prev<o.finallyLoc)return a(o.finallyLoc)}}}},abrupt:function(t,e){for(var n=this.tryEntries.length-1;n>=0;--n){var a=this.tryEntries[n];if(a.tryLoc<=this.prev&&r.call(a,"finallyLoc")&&this.prev<a.finallyLoc){var i=a;break}}i&&("break"===t||"continue"===t)&&i.tryLoc<=e&&e<=i.finallyLoc&&(i=null);var o=i?i.completion:{};return o.type=t,o.arg=e,i?(this.method="next",this.next=i.finallyLoc,g):this.complete(o)},complete:function(t,e){if("throw"===t.type)throw t.arg;return"break"===t.type||"continue"===t.type?this.next=t.arg:"return"===t.type?(this.rval=this.arg=t.arg,this.method="return",this.next="end"):"normal"===t.type&&e&&(this.next=e),g},finish:function(t){for(var e=this.tryEntries.length-1;e>=0;--e){var n=this.tryEntries[e];if(n.finallyLoc===t)return this.complete(n.completion,n.afterLoc),_(n),g}},catch:function(t){for(var e=this.tryEntries.length-1;e>=0;--e){var n=this.tryEntries[e];if(n.tryLoc===t){var r=n.completion;if("throw"===r.type){var a=r.arg;_(n)}return a}}throw new Error("illegal catch attempt")},delegateYield:function(t,n,r){return this.delegate={iterator:N(t),resultName:n,nextLoc:r},"next"===this.method&&(this.arg=e),g}},t}(t.exports);try{regeneratorRuntime=r}catch(a){Function("r","regeneratorRuntime = r")(r)}},a9e3:function(t,e,n){"use strict";var r=n("83ab"),a=n("da84"),i=n("94ca"),o=n("6eeb"),u=n("5135"),s=n("c6b6"),c=n("7156"),d=n("c04e"),f=n("d039"),l=n("7c73"),h=n("241c").f,m=n("06cf").f,g=n("9bf2").f,v=n("58a8").trim,w="Number",p=a[w],y=p.prototype,b=s(l(y))==w,T=function(t){var e,n,r,a,i,o,u,s,c=d(t,!1);if("string"==typeof c&&c.length>2)if(c=v(c),e=c.charCodeAt(0),43===e||45===e){if(n=c.charCodeAt(2),88===n||120===n)return NaN}else if(48===e){switch(c.charCodeAt(1)){case 66:case 98:r=2,a=49;break;case 79:case 111:r=8,a=55;break;default:return+c}for(i=c.slice(2),o=i.length,u=0;u<o;u++)if(s=i.charCodeAt(u),s<48||s>a)return NaN;return parseInt(i,r)}return+c};if(i(w,!p(" 0o1")||!p("0b1")||p("+0x1"))){for(var x,C=function(t){var e=arguments.length<1?0:t,n=this;return n instanceof C&&(b?f((function(){y.valueOf.call(n)})):s(n)!=w)?c(new p(T(e)),n,C):T(e)},M=r?h(p):"MAX_VALUE,MIN_VALUE,NaN,NEGATIVE_INFINITY,POSITIVE_INFINITY,EPSILON,isFinite,isInteger,isNaN,isSafeInteger,MAX_SAFE_INTEGER,MIN_SAFE_INTEGER,parseFloat,parseInt,isInteger,fromString,range".split(","),S=0;M.length>S;S++)u(p,x=M[S])&&!u(C,x)&&g(C,x,m(p,x));C.prototype=y,y.constructor=C,o(a,w,C)}},ac7f:function(t,e,n){"use strict";n.r(e);var r=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{attrs:{id:"player-page"}},[n("div",{attrs:{id:"side-panel"}},[n("h1",[t._v("Player")]),n("input",{attrs:{type:"file",value:"読み込み"},on:{change:t.loadData}}),n("VideoInfoComponent",{attrs:{videoInfo:t.videoInfo}})],1),n("div",{attrs:{id:"player-panel"}},[n("textarea",{attrs:{id:"editor-aria"}}),n("VideoSliderBar",{attrs:{elapsedTime:t.elapsedTime,totalTime:t.totalTime}})],1)])},a=[];n("d3b7");function i(t,e,n,r,a,i,o){try{var u=t[i](o),s=u.value}catch(c){return void n(c)}u.done?e(s):Promise.resolve(s).then(r,a)}function o(t){return function(){var e=this,n=arguments;return new Promise((function(r,a){var o=t.apply(e,n);function u(t){i(o,r,a,u,s,"next",t)}function s(t){i(o,r,a,u,s,"throw",t)}u(void 0)}))}}n("96cf");var u=n("2b0e"),s=(n("81c6"),n("56b3")),c=n.n(s),d=(n("a7be"),n("7a7a"),n("4ba6"),n("9b74"),n("f6b6"),n("f9d4"),n("3c98"),n("10b2"),function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{attrs:{id:"video-info"}},[n("p",[t._v(" タイトル:"),n("cite",[t._v(t._s(t.videoInfo.title))])]),n("p",[t._v(" 作者:"),n("cite",[t._v(t._s(t.videoInfo.name))])]),n("p",[t._v("言語:"+t._s(t.language))]),n("p",[t._v("投稿日:"+t._s(t.uploadTime))]),n("p",[t._v("動画時間:"+t._s(t.recordingTime))])])}),f=[];n("b0c0");function l(t,e){if(e.length<t)throw new TypeError(t+" argument"+(t>1?"s":"")+" required, but only "+e.length+" present")}function h(t){l(1,arguments);var e=Object.prototype.toString.call(t);return t instanceof Date||"object"===typeof t&&"[object Date]"===e?new Date(t.getTime()):"number"===typeof t||"[object Number]"===e?new Date(t):("string"!==typeof t&&"[object String]"!==e||"undefined"===typeof console||(console.warn("Starting with v2.0.0-beta.1 date-fns doesn't accept strings as date arguments. Please use `parseISO` to parse strings. See: https://git.io/fjule"),console.warn((new Error).stack)),new Date(NaN))}function m(t){l(1,arguments);var e=h(t);return!isNaN(e)}var g={lessThanXSeconds:{one:"less than a second",other:"less than {{count}} seconds"},xSeconds:{one:"1 second",other:"{{count}} seconds"},halfAMinute:"half a minute",lessThanXMinutes:{one:"less than a minute",other:"less than {{count}} minutes"},xMinutes:{one:"1 minute",other:"{{count}} minutes"},aboutXHours:{one:"about 1 hour",other:"about {{count}} hours"},xHours:{one:"1 hour",other:"{{count}} hours"},xDays:{one:"1 day",other:"{{count}} days"},aboutXWeeks:{one:"about 1 week",other:"about {{count}} weeks"},xWeeks:{one:"1 week",other:"{{count}} weeks"},aboutXMonths:{one:"about 1 month",other:"about {{count}} months"},xMonths:{one:"1 month",other:"{{count}} months"},aboutXYears:{one:"about 1 year",other:"about {{count}} years"},xYears:{one:"1 year",other:"{{count}} years"},overXYears:{one:"over 1 year",other:"over {{count}} years"},almostXYears:{one:"almost 1 year",other:"almost {{count}} years"}};function v(t,e,n){var r;return n=n||{},r="string"===typeof g[t]?g[t]:1===e?g[t].one:g[t].other.replace("{{count}}",e),n.addSuffix?n.comparison>0?"in "+r:r+" ago":r}function w(t){return function(e){var n=e||{},r=n.width?String(n.width):t.defaultWidth,a=t.formats[r]||t.formats[t.defaultWidth];return a}}var p={full:"EEEE, MMMM do, y",long:"MMMM do, y",medium:"MMM d, y",short:"MM/dd/yyyy"},y={full:"h:mm:ss a zzzz",long:"h:mm:ss a z",medium:"h:mm:ss a",short:"h:mm a"},b={full:"{{date}} 'at' {{time}}",long:"{{date}} 'at' {{time}}",medium:"{{date}}, {{time}}",short:"{{date}}, {{time}}"},T={date:w({formats:p,defaultWidth:"full"}),time:w({formats:y,defaultWidth:"full"}),dateTime:w({formats:b,defaultWidth:"full"})},x=T,C={lastWeek:"'last' eeee 'at' p",yesterday:"'yesterday at' p",today:"'today at' p",tomorrow:"'tomorrow at' p",nextWeek:"eeee 'at' p",other:"P"};function M(t,e,n,r){return C[t]}function S(t){return function(e,n){var r,a=n||{},i=a.context?String(a.context):"standalone";if("formatting"===i&&t.formattingValues){var o=t.defaultFormattingWidth||t.defaultWidth,u=a.width?String(a.width):o;r=t.formattingValues[u]||t.formattingValues[o]}else{var s=t.defaultWidth,c=a.width?String(a.width):t.defaultWidth;r=t.values[c]||t.values[s]}var d=t.argumentCallback?t.argumentCallback(e):e;return r[d]}}var E={narrow:["B","A"],abbreviated:["BC","AD"],wide:["Before Christ","Anno Domini"]},k={narrow:["1","2","3","4"],abbreviated:["Q1","Q2","Q3","Q4"],wide:["1st quarter","2nd quarter","3rd quarter","4th quarter"]},_={narrow:["J","F","M","A","M","J","J","A","S","O","N","D"],abbreviated:["Jan","Feb","Mar","Apr","May","Jun","Jul","Aug","Sep","Oct","Nov","Dec"],wide:["January","February","March","April","May","June","July","August","September","October","November","December"]},P={narrow:["S","M","T","W","T","F","S"],short:["Su","Mo","Tu","We","Th","Fr","Sa"],abbreviated:["Sun","Mon","Tue","Wed","Thu","Fri","Sat"],wide:["Sunday","Monday","Tuesday","Wednesday","Thursday","Friday","Saturday"]},N={narrow:{am:"a",pm:"p",midnight:"mi",noon:"n",morning:"morning",afternoon:"afternoon",evening:"evening",night:"night"},abbreviated:{am:"AM",pm:"PM",midnight:"midnight",noon:"noon",morning:"morning",afternoon:"afternoon",evening:"evening",night:"night"},wide:{am:"a.m.",pm:"p.m.",midnight:"midnight",noon:"noon",morning:"morning",afternoon:"afternoon",evening:"evening",night:"night"}},D={narrow:{am:"a",pm:"p",midnight:"mi",noon:"n",morning:"in the morning",afternoon:"in the afternoon",evening:"in the evening",night:"at night"},abbreviated:{am:"AM",pm:"PM",midnight:"midnight",noon:"noon",morning:"in the morning",afternoon:"in the afternoon",evening:"in the evening",night:"at night"},wide:{am:"a.m.",pm:"p.m.",midnight:"midnight",noon:"noon",morning:"in the morning",afternoon:"in the afternoon",evening:"in the evening",night:"at night"}};function U(t,e){var n=Number(t),r=n%100;if(r>20||r<10)switch(r%10){case 1:return n+"st";case 2:return n+"nd";case 3:return n+"rd"}return n+"th"}var O={ordinalNumber:U,era:S({values:E,defaultWidth:"wide"}),quarter:S({values:k,defaultWidth:"wide",argumentCallback:function(t){return Number(t)-1}}),month:S({values:_,defaultWidth:"wide"}),day:S({values:P,defaultWidth:"wide"}),dayPeriod:S({values:N,defaultWidth:"wide",formattingValues:D,defaultFormattingWidth:"wide"})},L=O;function I(t){return function(e,n){var r=String(e),a=n||{},i=r.match(t.matchPattern);if(!i)return null;var o=i[0],u=r.match(t.parsePattern);if(!u)return null;var s=t.valueCallback?t.valueCallback(u[0]):u[0];return s=a.valueCallback?a.valueCallback(s):s,{value:s,rest:r.slice(o.length)}}}function W(t){return function(e,n){var r=String(e),a=n||{},i=a.width,o=i&&t.matchPatterns[i]||t.matchPatterns[t.defaultMatchWidth],u=r.match(o);if(!u)return null;var s,c=u[0],d=i&&t.parsePatterns[i]||t.parsePatterns[t.defaultParseWidth];return s="[object Array]"===Object.prototype.toString.call(d)?j(d,(function(t){return t.test(c)})):Y(d,(function(t){return t.test(c)})),s=t.valueCallback?t.valueCallback(s):s,s=a.valueCallback?a.valueCallback(s):s,{value:s,rest:r.slice(c.length)}}}function Y(t,e){for(var n in t)if(t.hasOwnProperty(n)&&e(t[n]))return n}function j(t,e){for(var n=0;n<t.length;n++)if(e(t[n]))return n}var q=/^(\d+)(th|st|nd|rd)?/i,F=/\d+/i,A={narrow:/^(b|a)/i,abbreviated:/^(b\.?\s?c\.?|b\.?\s?c\.?\s?e\.?|a\.?\s?d\.?|c\.?\s?e\.?)/i,wide:/^(before christ|before common era|anno domini|common era)/i},G={any:[/^b/i,/^(a|c)/i]},H={narrow:/^[1234]/i,abbreviated:/^q[1234]/i,wide:/^[1234](th|st|nd|rd)? quarter/i},z={any:[/1/i,/2/i,/3/i,/4/i]},X={narrow:/^[jfmasond]/i,abbreviated:/^(jan|feb|mar|apr|may|jun|jul|aug|sep|oct|nov|dec)/i,wide:/^(january|february|march|april|may|june|july|august|september|october|november|december)/i},R={narrow:[/^j/i,/^f/i,/^m/i,/^a/i,/^m/i,/^j/i,/^j/i,/^a/i,/^s/i,/^o/i,/^n/i,/^d/i],any:[/^ja/i,/^f/i,/^mar/i,/^ap/i,/^may/i,/^jun/i,/^jul/i,/^au/i,/^s/i,/^o/i,/^n/i,/^d/i]},B={narrow:/^[smtwf]/i,short:/^(su|mo|tu|we|th|fr|sa)/i,abbreviated:/^(sun|mon|tue|wed|thu|fri|sat)/i,wide:/^(sunday|monday|tuesday|wednesday|thursday|friday|saturday)/i},Q={narrow:[/^s/i,/^m/i,/^t/i,/^w/i,/^t/i,/^f/i,/^s/i],any:[/^su/i,/^m/i,/^tu/i,/^w/i,/^th/i,/^f/i,/^sa/i]},V={narrow:/^(a|p|mi|n|(in the|at) (morning|afternoon|evening|night))/i,any:/^([ap]\.?\s?m\.?|midnight|noon|(in the|at) (morning|afternoon|evening|night))/i},J={any:{am:/^a/i,pm:/^p/i,midnight:/^mi/i,noon:/^no/i,morning:/morning/i,afternoon:/afternoon/i,evening:/evening/i,night:/night/i}},$={ordinalNumber:I({matchPattern:q,parsePattern:F,valueCallback:function(t){return parseInt(t,10)}}),era:W({matchPatterns:A,defaultMatchWidth:"wide",parsePatterns:G,defaultParseWidth:"any"}),quarter:W({matchPatterns:H,defaultMatchWidth:"wide",parsePatterns:z,defaultParseWidth:"any",valueCallback:function(t){return t+1}}),month:W({matchPatterns:X,defaultMatchWidth:"wide",parsePatterns:R,defaultParseWidth:"any"}),day:W({matchPatterns:B,defaultMatchWidth:"wide",parsePatterns:Q,defaultParseWidth:"any"}),dayPeriod:W({matchPatterns:V,defaultMatchWidth:"any",parsePatterns:J,defaultParseWidth:"any"})},K=$,Z={code:"en-US",formatDistance:v,formatLong:x,formatRelative:M,localize:L,match:K,options:{weekStartsOn:0,firstWeekContainsDate:1}},tt=Z;function et(t){if(null===t||!0===t||!1===t)return NaN;var e=Number(t);return isNaN(e)?e:e<0?Math.ceil(e):Math.floor(e)}function nt(t,e){l(2,arguments);var n=h(t).getTime(),r=et(e);return new Date(n+r)}function rt(t,e){l(2,arguments);var n=et(e);return nt(t,-n)}function at(t,e){var n=t<0?"-":"",r=Math.abs(t).toString();while(r.length<e)r="0"+r;return n+r}var it={y:function(t,e){var n=t.getUTCFullYear(),r=n>0?n:1-n;return at("yy"===e?r%100:r,e.length)},M:function(t,e){var n=t.getUTCMonth();return"M"===e?String(n+1):at(n+1,2)},d:function(t,e){return at(t.getUTCDate(),e.length)},a:function(t,e){var n=t.getUTCHours()/12>=1?"pm":"am";switch(e){case"a":case"aa":return n.toUpperCase();case"aaa":return n;case"aaaaa":return n[0];case"aaaa":default:return"am"===n?"a.m.":"p.m."}},h:function(t,e){return at(t.getUTCHours()%12||12,e.length)},H:function(t,e){return at(t.getUTCHours(),e.length)},m:function(t,e){return at(t.getUTCMinutes(),e.length)},s:function(t,e){return at(t.getUTCSeconds(),e.length)},S:function(t,e){var n=e.length,r=t.getUTCMilliseconds(),a=Math.floor(r*Math.pow(10,n-3));return at(a,e.length)}},ot=it,ut=864e5;function st(t){l(1,arguments);var e=h(t),n=e.getTime();e.setUTCMonth(0,1),e.setUTCHours(0,0,0,0);var r=e.getTime(),a=n-r;return Math.floor(a/ut)+1}function ct(t){l(1,arguments);var e=1,n=h(t),r=n.getUTCDay(),a=(r<e?7:0)+r-e;return n.setUTCDate(n.getUTCDate()-a),n.setUTCHours(0,0,0,0),n}function dt(t){l(1,arguments);var e=h(t),n=e.getUTCFullYear(),r=new Date(0);r.setUTCFullYear(n+1,0,4),r.setUTCHours(0,0,0,0);var a=ct(r),i=new Date(0);i.setUTCFullYear(n,0,4),i.setUTCHours(0,0,0,0);var o=ct(i);return e.getTime()>=a.getTime()?n+1:e.getTime()>=o.getTime()?n:n-1}function ft(t){l(1,arguments);var e=dt(t),n=new Date(0);n.setUTCFullYear(e,0,4),n.setUTCHours(0,0,0,0);var r=ct(n);return r}var lt=6048e5;function ht(t){l(1,arguments);var e=h(t),n=ct(e).getTime()-ft(e).getTime();return Math.round(n/lt)+1}function mt(t,e){l(1,arguments);var n=e||{},r=n.locale,a=r&&r.options&&r.options.weekStartsOn,i=null==a?0:et(a),o=null==n.weekStartsOn?i:et(n.weekStartsOn);if(!(o>=0&&o<=6))throw new RangeError("weekStartsOn must be between 0 and 6 inclusively");var u=h(t),s=u.getUTCDay(),c=(s<o?7:0)+s-o;return u.setUTCDate(u.getUTCDate()-c),u.setUTCHours(0,0,0,0),u}function gt(t,e){l(1,arguments);var n=h(t,e),r=n.getUTCFullYear(),a=e||{},i=a.locale,o=i&&i.options&&i.options.firstWeekContainsDate,u=null==o?1:et(o),s=null==a.firstWeekContainsDate?u:et(a.firstWeekContainsDate);if(!(s>=1&&s<=7))throw new RangeError("firstWeekContainsDate must be between 1 and 7 inclusively");var c=new Date(0);c.setUTCFullYear(r+1,0,s),c.setUTCHours(0,0,0,0);var d=mt(c,e),f=new Date(0);f.setUTCFullYear(r,0,s),f.setUTCHours(0,0,0,0);var m=mt(f,e);return n.getTime()>=d.getTime()?r+1:n.getTime()>=m.getTime()?r:r-1}function vt(t,e){l(1,arguments);var n=e||{},r=n.locale,a=r&&r.options&&r.options.firstWeekContainsDate,i=null==a?1:et(a),o=null==n.firstWeekContainsDate?i:et(n.firstWeekContainsDate),u=gt(t,e),s=new Date(0);s.setUTCFullYear(u,0,o),s.setUTCHours(0,0,0,0);var c=mt(s,e);return c}var wt=6048e5;function pt(t,e){l(1,arguments);var n=h(t),r=mt(n,e).getTime()-vt(n,e).getTime();return Math.round(r/wt)+1}var yt={am:"am",pm:"pm",midnight:"midnight",noon:"noon",morning:"morning",afternoon:"afternoon",evening:"evening",night:"night"},bt={G:function(t,e,n){var r=t.getUTCFullYear()>0?1:0;switch(e){case"G":case"GG":case"GGG":return n.era(r,{width:"abbreviated"});case"GGGGG":return n.era(r,{width:"narrow"});case"GGGG":default:return n.era(r,{width:"wide"})}},y:function(t,e,n){if("yo"===e){var r=t.getUTCFullYear(),a=r>0?r:1-r;return n.ordinalNumber(a,{unit:"year"})}return ot.y(t,e)},Y:function(t,e,n,r){var a=gt(t,r),i=a>0?a:1-a;if("YY"===e){var o=i%100;return at(o,2)}return"Yo"===e?n.ordinalNumber(i,{unit:"year"}):at(i,e.length)},R:function(t,e){var n=dt(t);return at(n,e.length)},u:function(t,e){var n=t.getUTCFullYear();return at(n,e.length)},Q:function(t,e,n){var r=Math.ceil((t.getUTCMonth()+1)/3);switch(e){case"Q":return String(r);case"QQ":return at(r,2);case"Qo":return n.ordinalNumber(r,{unit:"quarter"});case"QQQ":return n.quarter(r,{width:"abbreviated",context:"formatting"});case"QQQQQ":return n.quarter(r,{width:"narrow",context:"formatting"});case"QQQQ":default:return n.quarter(r,{width:"wide",context:"formatting"})}},q:function(t,e,n){var r=Math.ceil((t.getUTCMonth()+1)/3);switch(e){case"q":return String(r);case"qq":return at(r,2);case"qo":return n.ordinalNumber(r,{unit:"quarter"});case"qqq":return n.quarter(r,{width:"abbreviated",context:"standalone"});case"qqqqq":return n.quarter(r,{width:"narrow",context:"standalone"});case"qqqq":default:return n.quarter(r,{width:"wide",context:"standalone"})}},M:function(t,e,n){var r=t.getUTCMonth();switch(e){case"M":case"MM":return ot.M(t,e);case"Mo":return n.ordinalNumber(r+1,{unit:"month"});case"MMM":return n.month(r,{width:"abbreviated",context:"formatting"});case"MMMMM":return n.month(r,{width:"narrow",context:"formatting"});case"MMMM":default:return n.month(r,{width:"wide",context:"formatting"})}},L:function(t,e,n){var r=t.getUTCMonth();switch(e){case"L":return String(r+1);case"LL":return at(r+1,2);case"Lo":return n.ordinalNumber(r+1,{unit:"month"});case"LLL":return n.month(r,{width:"abbreviated",context:"standalone"});case"LLLLL":return n.month(r,{width:"narrow",context:"standalone"});case"LLLL":default:return n.month(r,{width:"wide",context:"standalone"})}},w:function(t,e,n,r){var a=pt(t,r);return"wo"===e?n.ordinalNumber(a,{unit:"week"}):at(a,e.length)},I:function(t,e,n){var r=ht(t);return"Io"===e?n.ordinalNumber(r,{unit:"week"}):at(r,e.length)},d:function(t,e,n){return"do"===e?n.ordinalNumber(t.getUTCDate(),{unit:"date"}):ot.d(t,e)},D:function(t,e,n){var r=st(t);return"Do"===e?n.ordinalNumber(r,{unit:"dayOfYear"}):at(r,e.length)},E:function(t,e,n){var r=t.getUTCDay();switch(e){case"E":case"EE":case"EEE":return n.day(r,{width:"abbreviated",context:"formatting"});case"EEEEE":return n.day(r,{width:"narrow",context:"formatting"});case"EEEEEE":return n.day(r,{width:"short",context:"formatting"});case"EEEE":default:return n.day(r,{width:"wide",context:"formatting"})}},e:function(t,e,n,r){var a=t.getUTCDay(),i=(a-r.weekStartsOn+8)%7||7;switch(e){case"e":return String(i);case"ee":return at(i,2);case"eo":return n.ordinalNumber(i,{unit:"day"});case"eee":return n.day(a,{width:"abbreviated",context:"formatting"});case"eeeee":return n.day(a,{width:"narrow",context:"formatting"});case"eeeeee":return n.day(a,{width:"short",context:"formatting"});case"eeee":default:return n.day(a,{width:"wide",context:"formatting"})}},c:function(t,e,n,r){var a=t.getUTCDay(),i=(a-r.weekStartsOn+8)%7||7;switch(e){case"c":return String(i);case"cc":return at(i,e.length);case"co":return n.ordinalNumber(i,{unit:"day"});case"ccc":return n.day(a,{width:"abbreviated",context:"standalone"});case"ccccc":return n.day(a,{width:"narrow",context:"standalone"});case"cccccc":return n.day(a,{width:"short",context:"standalone"});case"cccc":default:return n.day(a,{width:"wide",context:"standalone"})}},i:function(t,e,n){var r=t.getUTCDay(),a=0===r?7:r;switch(e){case"i":return String(a);case"ii":return at(a,e.length);case"io":return n.ordinalNumber(a,{unit:"day"});case"iii":return n.day(r,{width:"abbreviated",context:"formatting"});case"iiiii":return n.day(r,{width:"narrow",context:"formatting"});case"iiiiii":return n.day(r,{width:"short",context:"formatting"});case"iiii":default:return n.day(r,{width:"wide",context:"formatting"})}},a:function(t,e,n){var r=t.getUTCHours(),a=r/12>=1?"pm":"am";switch(e){case"a":case"aa":return n.dayPeriod(a,{width:"abbreviated",context:"formatting"});case"aaa":return n.dayPeriod(a,{width:"abbreviated",context:"formatting"}).toLowerCase();case"aaaaa":return n.dayPeriod(a,{width:"narrow",context:"formatting"});case"aaaa":default:return n.dayPeriod(a,{width:"wide",context:"formatting"})}},b:function(t,e,n){var r,a=t.getUTCHours();switch(r=12===a?yt.noon:0===a?yt.midnight:a/12>=1?"pm":"am",e){case"b":case"bb":return n.dayPeriod(r,{width:"abbreviated",context:"formatting"});case"bbb":return n.dayPeriod(r,{width:"abbreviated",context:"formatting"}).toLowerCase();case"bbbbb":return n.dayPeriod(r,{width:"narrow",context:"formatting"});case"bbbb":default:return n.dayPeriod(r,{width:"wide",context:"formatting"})}},B:function(t,e,n){var r,a=t.getUTCHours();switch(r=a>=17?yt.evening:a>=12?yt.afternoon:a>=4?yt.morning:yt.night,e){case"B":case"BB":case"BBB":return n.dayPeriod(r,{width:"abbreviated",context:"formatting"});case"BBBBB":return n.dayPeriod(r,{width:"narrow",context:"formatting"});case"BBBB":default:return n.dayPeriod(r,{width:"wide",context:"formatting"})}},h:function(t,e,n){if("ho"===e){var r=t.getUTCHours()%12;return 0===r&&(r=12),n.ordinalNumber(r,{unit:"hour"})}return ot.h(t,e)},H:function(t,e,n){return"Ho"===e?n.ordinalNumber(t.getUTCHours(),{unit:"hour"}):ot.H(t,e)},K:function(t,e,n){var r=t.getUTCHours()%12;return"Ko"===e?n.ordinalNumber(r,{unit:"hour"}):at(r,e.length)},k:function(t,e,n){var r=t.getUTCHours();return 0===r&&(r=24),"ko"===e?n.ordinalNumber(r,{unit:"hour"}):at(r,e.length)},m:function(t,e,n){return"mo"===e?n.ordinalNumber(t.getUTCMinutes(),{unit:"minute"}):ot.m(t,e)},s:function(t,e,n){return"so"===e?n.ordinalNumber(t.getUTCSeconds(),{unit:"second"}):ot.s(t,e)},S:function(t,e){return ot.S(t,e)},X:function(t,e,n,r){var a=r._originalDate||t,i=a.getTimezoneOffset();if(0===i)return"Z";switch(e){case"X":return xt(i);case"XXXX":case"XX":return Ct(i);case"XXXXX":case"XXX":default:return Ct(i,":")}},x:function(t,e,n,r){var a=r._originalDate||t,i=a.getTimezoneOffset();switch(e){case"x":return xt(i);case"xxxx":case"xx":return Ct(i);case"xxxxx":case"xxx":default:return Ct(i,":")}},O:function(t,e,n,r){var a=r._originalDate||t,i=a.getTimezoneOffset();switch(e){case"O":case"OO":case"OOO":return"GMT"+Tt(i,":");case"OOOO":default:return"GMT"+Ct(i,":")}},z:function(t,e,n,r){var a=r._originalDate||t,i=a.getTimezoneOffset();switch(e){case"z":case"zz":case"zzz":return"GMT"+Tt(i,":");case"zzzz":default:return"GMT"+Ct(i,":")}},t:function(t,e,n,r){var a=r._originalDate||t,i=Math.floor(a.getTime()/1e3);return at(i,e.length)},T:function(t,e,n,r){var a=r._originalDate||t,i=a.getTime();return at(i,e.length)}};function Tt(t,e){var n=t>0?"-":"+",r=Math.abs(t),a=Math.floor(r/60),i=r%60;if(0===i)return n+String(a);var o=e||"";return n+String(a)+o+at(i,2)}function xt(t,e){if(t%60===0){var n=t>0?"-":"+";return n+at(Math.abs(t)/60,2)}return Ct(t,e)}function Ct(t,e){var n=e||"",r=t>0?"-":"+",a=Math.abs(t),i=at(Math.floor(a/60),2),o=at(a%60,2);return r+i+n+o}var Mt=bt;function St(t,e){switch(t){case"P":return e.date({width:"short"});case"PP":return e.date({width:"medium"});case"PPP":return e.date({width:"long"});case"PPPP":default:return e.date({width:"full"})}}function Et(t,e){switch(t){case"p":return e.time({width:"short"});case"pp":return e.time({width:"medium"});case"ppp":return e.time({width:"long"});case"pppp":default:return e.time({width:"full"})}}function kt(t,e){var n,r=t.match(/(P+)(p+)?/),a=r[1],i=r[2];if(!i)return St(t,e);switch(a){case"P":n=e.dateTime({width:"short"});break;case"PP":n=e.dateTime({width:"medium"});break;case"PPP":n=e.dateTime({width:"long"});break;case"PPPP":default:n=e.dateTime({width:"full"});break}return n.replace("{{date}}",St(a,e)).replace("{{time}}",Et(i,e))}var _t={p:Et,P:kt},Pt=_t;function Nt(t){var e=new Date(Date.UTC(t.getFullYear(),t.getMonth(),t.getDate(),t.getHours(),t.getMinutes(),t.getSeconds(),t.getMilliseconds()));return e.setUTCFullYear(t.getFullYear()),t.getTime()-e.getTime()}var Dt=["D","DD"],Ut=["YY","YYYY"];function Ot(t){return-1!==Dt.indexOf(t)}function Lt(t){return-1!==Ut.indexOf(t)}function It(t,e,n){if("YYYY"===t)throw new RangeError("Use `yyyy` instead of `YYYY` (in `".concat(e,"`) for formatting years to the input `").concat(n,"`; see: https://git.io/fxCyr"));if("YY"===t)throw new RangeError("Use `yy` instead of `YY` (in `".concat(e,"`) for formatting years to the input `").concat(n,"`; see: https://git.io/fxCyr"));if("D"===t)throw new RangeError("Use `d` instead of `D` (in `".concat(e,"`) for formatting days of the month to the input `").concat(n,"`; see: https://git.io/fxCyr"));if("DD"===t)throw new RangeError("Use `dd` instead of `DD` (in `".concat(e,"`) for formatting days of the month to the input `").concat(n,"`; see: https://git.io/fxCyr"))}var Wt=/[yYQqMLwIdDecihHKkms]o|(\w)\1*|''|'(''|[^'])+('|$)|./g,Yt=/P+p+|P+|p+|''|'(''|[^'])+('|$)|./g,jt=/^'([^]*?)'?$/,qt=/''/g,Ft=/[a-zA-Z]/;function At(t,e,n){l(2,arguments);var r=String(e),a=n||{},i=a.locale||tt,o=i.options&&i.options.firstWeekContainsDate,u=null==o?1:et(o),s=null==a.firstWeekContainsDate?u:et(a.firstWeekContainsDate);if(!(s>=1&&s<=7))throw new RangeError("firstWeekContainsDate must be between 1 and 7 inclusively");var c=i.options&&i.options.weekStartsOn,d=null==c?0:et(c),f=null==a.weekStartsOn?d:et(a.weekStartsOn);if(!(f>=0&&f<=6))throw new RangeError("weekStartsOn must be between 0 and 6 inclusively");if(!i.localize)throw new RangeError("locale must contain localize property");if(!i.formatLong)throw new RangeError("locale must contain formatLong property");var g=h(t);if(!m(g))throw new RangeError("Invalid time value");var v=Nt(g),w=rt(g,v),p={firstWeekContainsDate:s,weekStartsOn:f,locale:i,_originalDate:g},y=r.match(Yt).map((function(t){var e=t[0];if("p"===e||"P"===e){var n=Pt[e];return n(t,i.formatLong,p)}return t})).join("").match(Wt).map((function(n){if("''"===n)return"'";var r=n[0];if("'"===r)return Gt(n);var o=Mt[r];if(o)return!a.useAdditionalWeekYearTokens&&Lt(n)&&It(n,e,t),!a.useAdditionalDayOfYearTokens&&Ot(n)&&It(n,e,t),o(w,n,i.localize,p);if(r.match(Ft))throw new RangeError("Format string contains an unescaped latin alphabet character `"+r+"`");return n})).join("");return y}function Gt(t){return t.match(jt)[1].replace(qt,"'")}var Ht=u["a"].extend({name:"VideoInfo",props:{videoInfo:{type:Object,default:function(){return{userId:-1,title:"",name:"",language:void 0,uploadTime:-1,recordingTime:-1}},required:!1}},computed:{language:function(){var t,e;return null!==(t=null===(e=this.videoInfo.language)||void 0===e?void 0:e.name)&&void 0!==t?t:""},uploadTime:function(){return this.videoInfo.uploadTime<0?"":At(this.videoInfo.uploadTime,"yyyy/MM/dd hh:mm:ss")},recordingTime:function(){return this.videoInfo.recordingTime<0?"":At(this.videoInfo.recordingTime,"mm分ss秒")}}}),zt=Ht,Xt=(n("3f30"),n("2877")),Rt=Object(Xt["a"])(zt,d,f,!1,null,"0d22ca70",null),Bt=Rt.exports,Qt=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"slider-component"},[n("div",{staticClass:"slider"},[n("div",{staticClass:"slider-value",style:t.sliderValueStyle})]),n("span",{staticClass:"elapsedTime"},[t._v(t._s(t.playbackPosition))])])},Vt=[],Jt=(n("a9e3"),n("99af"),u["a"].extend({name:"VideoSliderBar",props:{elapsedTime:{type:Number,default:0},totalTime:{type:Number,default:0}},computed:{sliderValueStyle:function(){return 0==this.totalTime?{width:"0%"}:{width:"".concat(this.elapsedTime/this.totalTime*100,"%")}},playbackPosition:function(){return"".concat(At(this.elapsedTime,"mm:ss"),"/").concat(At(this.totalTime,"mm:ss"))}}})),$t=Jt,Kt=(n("0d65"),Object(Xt["a"])($t,Qt,Vt,!1,null,"046580ea",null)),Zt=Kt.exports,te=n("d4ec"),ee=n("bee2"),ne=function(){function t(e){Object(te["a"])(this,t),this._index=0,this._codingSequence=e.value.concat()}return Object(ee["a"])(t,[{key:"isNext",value:function(){var t=this._index+1;return t!==this._codingSequence.length}},{key:"next",value:function(){return!!this.isNext()&&(this._index++,!0)}},{key:"from",get:function(){var t=this._index-1;if(-1!==t)return this._codingSequence[t]}},{key:"current",get:function(){var t=this._index;return this._codingSequence[t]}},{key:"to",get:function(){var t=this._index+1;if(t!==this._codingSequence.length)return this._codingSequence[t]}}]),t}(),re=function(){function t(){Object(te["a"])(this,t)}return Object(ee["a"])(t,[{key:"load",value:function(t,e){if(null==e)throw new Error("editor is undefined");this._videoInfo=t.header,this._stream=new ne(t);var n=this._videoInfo.language;if(void 0==n)throw new Error("video is not language data");e.setOption("mode",n.tag),e.setValue(""),e.focus()}},{key:"start",value:function(t,e){var n=this;if(null==t)throw new Error("editor is undefined");if(void 0==this._stream)throw new Error("video is not Load");ae((function(){if(void 0==n._stream)throw new Error("video is not Load");var r=n._stream.current.changeData,a=r.text,i=r.from,o=r.to,u=r.origin,s=n._stream.current.cursor;t.replaceRange(a,i,o,u),t.setCursor(s),e(n._stream),n._stream.next();var c=n._stream.isNext();if(void 0===n._stream.from)return e(n._stream),{isNext:c,nextSpan:n._stream.current.timestamp};if(void 0===n._stream.to){console.log("終了");var d=n._stream.current.changeData,f=d.text,l=d.from,h=d.to,m=d.origin,g=n._stream.current.cursor;return t.replaceRange(f,l,h,m),t.setCursor(g),e(n._stream),{isNext:c,nextSpan:1}}return e(n._stream),{isNext:c,nextSpan:n._stream.to.timestamp-n._stream.current.timestamp}}))}},{key:"videoInfo",get:function(){if(void 0==this._videoInfo)throw new Error("video is not Load");return this._videoInfo}}]),t}();function ae(t){var e=t(),n=e.isNext,r=e.nextSpan;n&&setTimeout((function(){ae(t)}),r)}function ie(t){return new Promise((function(e,n){var r=new FileReader;r.onload=function(){e(r.result)},r.onerror=n,r.readAsText(t)}))}var oe=u["a"].extend({name:"PlayerPage",components:{VideoInfoComponent:Bt,VideoSliderBar:Zt},data:function(){return{defualtConfig:{mode:"javascript",lineNumbers:!0,indentUnit:4,theme:"monokai",showHint:!0,readOnly:!0},player:new re,videoInfo:void 0,elapsedTime:0,totalTime:0}},methods:{loadData:function(){var t=o(regeneratorRuntime.mark((function t(e){var n,r,a,i,o,u,s=this;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:if(r=e.target,a=null===(n=r.files)||void 0===n?void 0:n.item(0),void 0!=a&&null!=a){t.next=4;break}return t.abrupt("return");case 4:return t.next=6,ie(a);case 6:i=t.sent,o=JSON.parse(i),this.player.load(o,this.editor),this.videoInfo=this.player.videoInfo,u=o.header.recordingTime,this.totalTime=u,this.player.start(this.editor,(function(t){t.isNext()?s.elapsedTime=t.current.timestamp:s.elapsedTime=s.totalTime}));case 13:case"end":return t.stop()}}),t,this)})));function e(e){return t.apply(this,arguments)}return e}()},mounted:function(){var t,e=document.querySelector("#editor-aria");if(null==e)throw new Error("textarea not found for CodeMirror");var n=this.defualtConfig;this.editor=c.a.fromTextArea(e,n),null===(t=this.editor)||void 0===t||t.setSize("100%","70vh")}}),ue=oe,se=(n("4f60"),Object(Xt["a"])(ue,r,a,!1,null,"5ff6c8f8",null));e["default"]=se.exports},b0c0:function(t,e,n){var r=n("83ab"),a=n("9bf2").f,i=Function.prototype,o=i.toString,u=/^\s*function ([^ (]*)/,s="name";r&&!(s in i)&&a(i,s,{configurable:!0,get:function(){try{return o.call(this).match(u)[1]}catch(t){return""}}})}}]);
//# sourceMappingURL=chunk-64f9a27d.e7188405.js.map