(function(e){function t(t){for(var n,s,c=t[0],i=t[1],u=t[2],l=0,p=[];l<c.length;l++)s=c[l],Object.prototype.hasOwnProperty.call(o,s)&&o[s]&&p.push(o[s][0]),o[s]=0;for(n in i)Object.prototype.hasOwnProperty.call(i,n)&&(e[n]=i[n]);f&&f(t);while(p.length)p.shift()();return a.push.apply(a,u||[]),r()}function r(){for(var e,t=0;t<a.length;t++){for(var r=a[t],n=!0,s=1;s<r.length;s++){var c=r[s];0!==o[c]&&(n=!1)}n&&(a.splice(t--,1),e=i(i.s=r[0]))}return e}var n={},s={app:0},o={app:0},a=[];function c(e){return i.p+"js/"+({Process:"Process","process-viewer":"process-viewer","test-page":"test-page"}[e]||e)+"."+{Process:"65060cd3","process-viewer":"d7bc17c6","test-page":"3d3df6b9"}[e]+".js"}function i(t){if(n[t])return n[t].exports;var r=n[t]={i:t,l:!1,exports:{}};return e[t].call(r.exports,r,r.exports,i),r.l=!0,r.exports}i.e=function(e){var t=[],r={Process:1,"process-viewer":1,"test-page":1};s[e]?t.push(s[e]):0!==s[e]&&r[e]&&t.push(s[e]=new Promise((function(t,r){for(var n="css/"+({Process:"Process","process-viewer":"process-viewer","test-page":"test-page"}[e]||e)+"."+{Process:"e54f6c9a","process-viewer":"c95703e2","test-page":"9f428ad4"}[e]+".css",o=i.p+n,a=document.getElementsByTagName("link"),c=0;c<a.length;c++){var u=a[c],l=u.getAttribute("data-href")||u.getAttribute("href");if("stylesheet"===u.rel&&(l===n||l===o))return t()}var p=document.getElementsByTagName("style");for(c=0;c<p.length;c++){u=p[c],l=u.getAttribute("data-href");if(l===n||l===o)return t()}var f=document.createElement("link");f.rel="stylesheet",f.type="text/css",f.onload=t,f.onerror=function(t){var n=t&&t.target&&t.target.src||o,a=new Error("Loading CSS chunk "+e+" failed.\n("+n+")");a.code="CSS_CHUNK_LOAD_FAILED",a.request=n,delete s[e],f.parentNode.removeChild(f),r(a)},f.href=o;var d=document.getElementsByTagName("head")[0];d.appendChild(f)})).then((function(){s[e]=0})));var n=o[e];if(0!==n)if(n)t.push(n[2]);else{var a=new Promise((function(t,r){n=o[e]=[t,r]}));t.push(n[2]=a);var u,l=document.createElement("script");l.charset="utf-8",l.timeout=120,i.nc&&l.setAttribute("nonce",i.nc),l.src=c(e);var p=new Error;u=function(t){l.onerror=l.onload=null,clearTimeout(f);var r=o[e];if(0!==r){if(r){var n=t&&("load"===t.type?"missing":t.type),s=t&&t.target&&t.target.src;p.message="Loading chunk "+e+" failed.\n("+n+": "+s+")",p.name="ChunkLoadError",p.type=n,p.request=s,r[1](p)}o[e]=void 0}};var f=setTimeout((function(){u({type:"timeout",target:l})}),12e4);l.onerror=l.onload=u,document.head.appendChild(l)}return Promise.all(t)},i.m=e,i.c=n,i.d=function(e,t,r){i.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:r})},i.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},i.t=function(e,t){if(1&t&&(e=i(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var r=Object.create(null);if(i.r(r),Object.defineProperty(r,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var n in e)i.d(r,n,function(t){return e[t]}.bind(null,n));return r},i.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return i.d(t,"a",t),t},i.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},i.p="/",i.oe=function(e){throw console.error(e),e};var u=window["webpackJsonp"]=window["webpackJsonp"]||[],l=u.push.bind(u);u.push=t,u=u.slice();for(var p=0;p<u.length;p++)t(u[p]);var f=l;a.push([0,"chunk-vendors"]),r()})({0:function(e,t,r){e.exports=r("56d7")},"034f":function(e,t,r){"use strict";r("85ec")},"56d7":function(e,t,r){"use strict";r.r(t);r("e260"),r("e6cf"),r("cca6"),r("a79d");var n=r("2b0e"),s=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",{staticClass:"app"},[r("nav",{staticClass:"nav"},[r("ul",{staticClass:"menu"},[r("li",{staticClass:"item"},[r("router-link",{staticClass:"nav-link",attrs:{to:"/"}},[e._v("Ended processes")])],1),r("li",{staticClass:"item"},[r("router-link",{staticClass:"nav-link",attrs:{to:"/process"}},[e._v("Process control")])],1),r("li",{staticClass:"item"},[r("router-link",{staticClass:"nav-link",attrs:{to:"/test-page"}},[e._v("Test page")])],1)])]),r("router-view",{staticClass:"content"})],1)},o=[],a=(r("034f"),r("2877")),c={},i=Object(a["a"])(c,s,o,!1,null,null,null),u=i.exports,l=(r("d3b7"),r("3ca3"),r("ddb0"),r("8c4f")),p=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",[r("h1",[e._v("Processes")]),e._l(e.processes,(function(t){return r("li",{key:t},[r("router-link",{attrs:{to:{name:"EndedProcessViewer",params:{processName:t}}}},[e._v(" "+e._s(t))])],1)}))],2)},f=[],d={name:"EndedProcesses",data:function(){return{processes:[]}},mounted:function(){var e=this;fetch("/api/processes").then((function(e){return e.json()})).then((function(t){console.log(t),e.processes=t["value"]}))}},v=d,h=Object(a["a"])(v,p,f,!1,null,null,null),m=h.exports;n["a"].use(l["a"]);var g=[{path:"/",name:"EndedProcesses",component:m},{path:"/process",name:"Process",component:function(){return r.e("Process").then(r.bind(null,"da71"))}},{path:"/test-page",name:"TestPage",component:function(){return r.e("test-page").then(r.bind(null,"7454"))}},{path:"/process-viewer/:processName",name:"EndedProcessViewer",component:function(){return r.e("process-viewer").then(r.bind(null,"d483"))},props:!0}],b=new l["a"]({mode:"history",base:"/",routes:g}),y=b,w=r("3746"),P=r.n(w);r("6c2c");n["a"].config.productionTip=!1,n["a"].use(P.a),new n["a"]({router:y,render:function(e){return e(u)}}).$mount("#app")},"85ec":function(e,t,r){}});
//# sourceMappingURL=app.52c61e01.js.map