(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["about"],{da71:function(t,e,r){"use strict";r.r(e);var n=r("7a23"),o={class:"process"};function c(t,e,r,c,s,a){var u=Object(n["u"])("ProcessControl");return Object(n["q"])(),Object(n["e"])("div",o,[Object(n["h"])(u)])}var s=Object(n["f"])("h1",null,"Photo processing",-1),a={key:0},u={key:1};function i(t,e,r,o,c,i){return Object(n["q"])(),Object(n["e"])("div",null,[s,c.error?(Object(n["q"])(),Object(n["e"])("div",a,Object(n["w"])(c.error),1)):Object(n["d"])("",!0),"working"!=c.motorStatus?(Object(n["q"])(),Object(n["e"])("form",u,[Object(n["f"])("button",{type:"button",onClick:e[0]||(e[0]=function(){return i.submit&&i.submit.apply(i,arguments)})},"START")])):Object(n["d"])("",!0),"working"==c.motorStatus?(Object(n["q"])(),Object(n["e"])("button",{key:2,onClick:e[1]||(e[1]=function(){return i.stopProcess&&i.stopProcess.apply(i,arguments)})}," Stop process ")):Object(n["d"])("",!0),Object(n["f"])("h4",null,"Motor status: "+Object(n["w"])(c.motorStatus),1)])}var p=r("1da1"),f=(r("96cf"),r("d3b7"),{name:"ProcessControl",data:function(){return{degrees:0,error:"",motorStatus:""}},methods:{stopProcess:function(){var t=this;fetch("/api/stop-process",{method:"POST",headers:{"Content-Type":"application/json"}}).then(function(){var t=Object(p["a"])(regeneratorRuntime.mark((function t(e){var r;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:if(e.ok){t.next=5;break}return t.next=3,e.text();case 3:throw r=t.sent,Error(r);case 5:case"end":return t.stop()}}),t)})));return function(e){return t.apply(this,arguments)}}()).catch((function(e){t.error=e}))},submit:function(){var t=this;fetch("/api/start-process",{method:"POST",headers:{"Content-Type":"application/json"}}).then(function(){var t=Object(p["a"])(regeneratorRuntime.mark((function t(e){var r;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:if(e.ok){t.next=5;break}return t.next=3,e.text();case 3:throw r=t.sent,Error(r);case 5:case"end":return t.stop()}}),t)})));return function(e){return t.apply(this,arguments)}}()).catch((function(e){t.error=e}))},requestStatus:function(){var t=this;fetch("/api/machine-status").then(function(){var t=Object(p["a"])(regeneratorRuntime.mark((function t(e){var r;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:if(e.ok){t.next=5;break}return t.next=3,e.text();case 3:throw r=t.sent,Error(r);case 5:return t.abrupt("return",e.json());case 6:case"end":return t.stop()}}),t)})));return function(e){return t.apply(this,arguments)}}()).then((function(e){t.motorStatus=e["value"]})).catch((function(e){t.error=e}))}},mounted:function(){setInterval(this.requestStatus,1e3)}});f.render=i;var b=f,h={name:"Process",components:{ProcessControl:b}};h.render=c;e["default"]=h}}]);
//# sourceMappingURL=about.509156ad.js.map