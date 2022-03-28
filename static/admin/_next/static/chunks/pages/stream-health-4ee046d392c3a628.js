(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[632],{24019:function(e,t,n){"use strict";n.d(t,{Z:function(){return l}});var r=n(1413),i=n(67294),a={icon:{tag:"svg",attrs:{viewBox:"64 64 896 896",focusable:"false"},children:[{tag:"path",attrs:{d:"M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448 448-200.6 448-448S759.4 64 512 64zm0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372 372 166.6 372 372-166.6 372-372 372z"}},{tag:"path",attrs:{d:"M686.7 638.6L544.1 535.5V288c0-4.4-3.6-8-8-8H488c-4.4 0-8 3.6-8 8v275.4c0 2.6 1.2 5 3.3 6.5l165.4 120.6c3.6 2.6 8.6 1.8 11.2-1.7l28.6-39c2.6-3.7 1.8-8.7-1.8-11.2z"}}]},name:"clock-circle",theme:"outlined"},o=n(42135),s=function(e,t){return i.createElement(o.Z,(0,r.Z)((0,r.Z)({},e),{},{ref:t,icon:a}))};s.displayName="ClockCircleOutlined";var l=i.forwardRef(s)},24308:function(e,t,n){"use strict";n.d(t,{c4:function(){return a}});var r=n(4942),i=n(87462),a=["xxl","xl","lg","md","sm","xs"],o={xs:"(max-width: 575px)",sm:"(min-width: 576px)",md:"(min-width: 768px)",lg:"(min-width: 992px)",xl:"(min-width: 1200px)",xxl:"(min-width: 1600px)"},s=new Map,l=-1,c={},u={matchHandlers:{},dispatch:function(e){return c=e,s.forEach((function(e){return e(c)})),s.size>=1},subscribe:function(e){return s.size||this.register(),l+=1,s.set(l,e),e(c),l},unsubscribe:function(e){s.delete(e),s.size||this.unregister()},unregister:function(){var e=this;Object.keys(o).forEach((function(t){var n=o[t],r=e.matchHandlers[n];null===r||void 0===r||r.mql.removeListener(null===r||void 0===r?void 0:r.listener)})),s.clear()},register:function(){var e=this;Object.keys(o).forEach((function(t){var n=o[t],a=function(n){var a=n.matches;e.dispatch((0,i.Z)((0,i.Z)({},c),(0,r.Z)({},t,a)))},s=window.matchMedia(n);s.addListener(a),e.matchHandlers[n]={mql:s,listener:a},a(s)}))}};t.ZP=u},27561:function(e,t,n){var r=n(67990),i=/^\s+/;e.exports=function(e){return e?e.slice(0,r(e)+1).replace(i,""):e}},67990:function(e){var t=/\s/;e.exports=function(e){for(var n=e.length;n--&&t.test(e.charAt(n)););return n}},33448:function(e,t,n){var r=n(44239),i=n(37005);e.exports=function(e){return"symbol"==typeof e||i(e)&&"[object Symbol]"==r(e)}},14841:function(e,t,n){var r=n(27561),i=n(13218),a=n(33448),o=/^[-+]0x[0-9a-f]+$/i,s=/^0b[01]+$/i,l=/^0o[0-7]+$/i,c=parseInt;e.exports=function(e){if("number"==typeof e)return e;if(a(e))return NaN;if(i(e)){var t="function"==typeof e.valueOf?e.valueOf():e;e=i(t)?t+"":t}if("string"!=typeof e)return 0===e?e:+e;e=r(e);var n=s.test(e);return n||l.test(e)?c(e.slice(2),n?2:8):o.test(e)?NaN:+e}},86401:function(e,t,n){(window.__NEXT_P=window.__NEXT_P||[]).push(["/stream-health",function(){return n(26102)}])},89270:function(e,t,n){"use strict";n.d(t,{Z:function(){return u}});var r=n(85893),i=n(31877),a=n(92616),o=n.n(a),s=n(58091),l=n(60727);function c(e){var t={};return e.forEach((function(e){var n=new Date(e.time),r=(0,s.Z)(n,"H:mma");t[r]=e.value})),t}function u(e){var t=e.data,n=e.title,i=e.color,a=e.unit,o=e.dataCollections,s=e.yFlipped,u=e.yLogarithmic,d=[];t&&t.length>0&&d.push({name:n,color:i,data:c(t)}),o.forEach((function(e){d.push({name:e.name,data:c(e.data),color:e.color,dataset:e.options})}));var f={scales:{y:{reverse:!1,type:"linear"},x:{type:"time"}}};return f.scales.y.reverse=s,f.scales.y.type=u?"logarithmic":"linear",(0,r.jsx)("div",{className:"line-chart-container",children:(0,r.jsx)(l.wW,{xtitle:"Time",ytitle:n,suffix:a,legend:"bottom",color:i,data:d,download:n,library:f})})}o().use(i.Z),u.defaultProps={dataCollections:[],data:[],title:"",yFlipped:!1,yLogarithmic:!1}},26102:function(e,t,n){"use strict";n.r(t),n.d(t,{default:function(){return C}});var r=n(34051),i=n.n(r),a=n(85893),o=n(84485),s=n(26713),l=n(25968),c=n(6226),u=n(97751),d=n(74763),f=n(14670),p=n(67294),h=n(1413),m={icon:{tag:"svg",attrs:{viewBox:"64 64 896 896",focusable:"false"},children:[{tag:"path",attrs:{d:"M723 620.5C666.8 571.6 593.4 542 513 542s-153.8 29.6-210.1 78.6a8.1 8.1 0 00-.8 11.2l36 42.9c2.9 3.4 8 3.8 11.4.9C393.1 637.2 450.3 614 513 614s119.9 23.2 163.5 61.5c3.4 2.9 8.5 2.5 11.4-.9l36-42.9c2.8-3.3 2.4-8.3-.9-11.2zm117.4-140.1C751.7 406.5 637.6 362 513 362s-238.7 44.5-327.5 118.4a8.05 8.05 0 00-1 11.3l36 42.9c2.8 3.4 7.9 3.8 11.2 1C308 472.2 406.1 434 513 434s205 38.2 281.2 101.6c3.4 2.8 8.4 2.4 11.2-1l36-42.9c2.8-3.4 2.4-8.5-1-11.3zm116.7-139C835.7 241.8 680.3 182 511 182c-168.2 0-322.6 59-443.7 157.4a8 8 0 00-1.1 11.4l36 42.9c2.8 3.3 7.8 3.8 11.1 1.1C222 306.7 360.3 254 511 254c151.8 0 291 53.5 400 142.7 3.4 2.8 8.4 2.3 11.2-1.1l36-42.9c2.9-3.4 2.4-8.5-1.1-11.3zM448 778a64 64 0 10128 0 64 64 0 10-128 0z"}}]},name:"wifi",theme:"outlined"},v=n(42135),y=function(e,t){return p.createElement(v.Z,(0,h.Z)((0,h.Z)({},e),{},{ref:t,icon:m}))};y.displayName="WifiOutlined";var g=p.forwardRef(y),w=n(24019),x=n(28058),b=n(58827),j=n(89270);function Z(e,t,n,r,i,a,o){try{var s=e[a](o),l=s.value}catch(c){return void n(c)}s.done?t(l):Promise.resolve(l).then(r,i)}function F(e){var t=e.title,n=e.description;return(0,a.jsxs)("div",{className:"description-box",children:[(0,a.jsx)(o.Z.Title,{children:t}),(0,a.jsx)(o.Z.Paragraph,{children:n})]})}function C(){var e,t,n,r,h,m,v,y=(0,p.useState)([]),C=y[0],k=y[1],P=(0,p.useState)([]),S=P[0],N=P[1],E=(0,p.useState)(),O=E[0],L=E[1],D=(0,p.useState)(),M=D[0],T=D[1],A=(0,p.useState)([]),q=A[0],I=A[1],z=(0,p.useState)([]),B=z[0],_=z[1],R=(0,p.useState)([]),V=R[0],H=R[1],W=(0,p.useState)([]),Q=W[0],$=W[1],X=(0,p.useState)([]),Y=X[0],G=X[1],J=(0,p.useState)([]),K=J[0],U=J[1],ee=(0,p.useState)([]),te=ee[0],ne=ee[1],re=(0,p.useState)([]),ie=re[0],ae=re[1],oe=(0,p.useState)(0),se=oe[0],le=oe[1],ce=(0,p.useState)(0),ue=ce[0],de=ce[1],fe=function(){var e,t=(e=i().mark((function e(){var t;return i().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.prev=0,e.next=3,(0,b.rQ)(b.N$);case 3:t=e.sent,k(t.errors),N(t.qualityVariantChanges),T(t.highestLatency),L(t.lowestLatency),I(t.medianLatency),_(t.medianSegmentDownloadDuration),H(t.maximumSegmentDownloadDuration),$(t.minimumSegmentDownloadDuration),G(t.minPlayerBitrate),U(t.medianPlayerBitrate),ne(t.maxPlayerBitrate),ae(t.availableBitrates),le(t.segmentLength-.3),de(t.representation),e.next=23;break;case 20:e.prev=20,e.t0=e.catch(0),console.error(e.t0);case 23:case"end":return e.stop()}}),e,null,[[0,20]])})),function(){var t=this,n=arguments;return new Promise((function(r,i){var a=e.apply(t,n);function o(e){Z(a,r,i,o,s,"next",e)}function s(e){Z(a,r,i,o,s,"throw",e)}o(void 0)}))});return function(){return t.apply(this,arguments)}}();(0,p.useEffect)((function(){var e;return fe(),e=setInterval(fe,b.NE),function(){clearInterval(e)}}),[]);var pe=(0,a.jsxs)("div",{children:[(0,a.jsx)(o.Z.Title,{children:"Stream Performance"}),(0,a.jsx)(o.Z.Paragraph,{children:"Data has not yet been collected. Once a stream has begun and viewers are watching this page will be available."})]});if(!(null===C||void 0===C?void 0:C.length))return pe;if(!(null===q||void 0===q?void 0:q.length))return pe;if(!(null===B||void 0===B?void 0:B.length))return pe;var he=[{name:"Errors",color:"#B63FFF",options:{radius:3},data:C},{name:"Quality changes",color:"#2087E2",options:{radius:2},data:S}],me=[{name:"Median stream latency",color:"#00FFFF",options:{radius:2},data:q},{name:"Lowest stream latency",color:"#02FD0D",options:{radius:2},data:O},{name:"Highest stream latency",color:"#B63FFF",options:{radius:2},data:M}],ve=[{name:"Max download duration",color:"#B63FFF",options:{radius:2},data:V},{name:"Median download duration",color:"#00FFFF",options:{radius:2},data:B},{name:"Min download duration",color:"#02FD0D",options:{radius:2},data:Q},{name:"Approximate limit",color:"#003FFF",data:B.map((function(e){return{time:e.time,value:se}})),options:{radius:0}}],ye=[{name:"Lowest viewer bitrate",color:"#B63FFF",data:Y,options:{radius:2}},{name:"Median viewer bitrate",color:"#00FFFF",data:K,options:{radius:2}},{name:"Maximum viewer bitrate",color:"#02FD0D",data:te,options:{radius:2}}];ie.forEach((function(e){ye.push({name:"Available bitrate",color:"#003FFF",data:Y.map((function(t){return{time:t.time,value:e}})),options:{radius:0}})}));var ge=null===(t=null===(e=ye[0])||void 0===e?void 0:e.data[ye[0].data.length-1])||void 0===t?void 0:t.value,we=null===(n=B[B.length-1])||void 0===n?void 0:n.value,xe=ie[0],be=(null===(r=q[q.length-1])||void 0===r?void 0:r.value)||0,je=(null===(h=M[M.length-1])||void 0===h?void 0:h.value)||0,Ze=(null===(m=O[O.length-1])||void 0===m?void 0:m.value)||0,Fe=(Number(je)+Number(Ze)+Number(be))/3,Ce=0;((null===(v=he[0])||void 0===v?void 0:v.data.length)||0)>5?Ce=he[0].data.slice(-5).reduce((function(e,t){return e+Number(t.value)}),0):Ce=he[0].data.reduce((function(e,t){return e+Number(t.value)}),0);var ke=ge>0||we>0||Ce>0,Pe=null,Se=null;0!==ge&&ge<xe&&(Pe="One of your viewers is playing your stream at ".concat(ge,"kbps, slower than ").concat(xe,"kbps, the lowest quality you made available. Consider adding a lower quality with a lower bitrate if the errors over time warrant this.")),we>se&&(Se="Your viewers may be consuming your video slower than required. This may be due to slow networks or your latency configuration. You need to decrease the amount of time viewers are taking to consume your video. Consider adding a lower quality with a lower bitrate or experiment with increasing the latency buffer setting.");var Ne=Ce>0?"#B63FFF":"#FFFFFF",Ee={display:"flex",alignItems:"center",justifyContent:"center",height:"80px"};return(0,a.jsxs)(a.Fragment,{children:[(0,a.jsx)(o.Z.Title,{children:"Stream Performance"}),(0,a.jsx)(o.Z.Paragraph,{children:"This tool hopes to help you identify and troubleshoot problems you may be experiencing with your stream. It aims to aggregate experiences across your viewers, meaning one viewer with an exceptionally bad experience may throw off numbers for the whole, especially with a low number of viewers."}),(0,a.jsx)(o.Z.Paragraph,{children:"The data is only collected by those using the Owncast web interface and is unable to gain insight into external players people may be using such as VLC, MPV, QuickTime, etc."}),(0,a.jsxs)(s.Z,{direction:"vertical",size:"middle",children:[(0,a.jsxs)(l.Z,{gutter:[16,16],justify:"space-around",style:{display:ke?"flex":"none"},children:[(0,a.jsx)(c.Z,{children:(0,a.jsx)(u.Z,{type:"inner",children:(0,a.jsx)("div",{style:Ee,children:(0,a.jsx)(d.Z,{title:"Viewer Playback Speed",value:"".concat(ge),prefix:(0,a.jsx)(g,{style:{marginRight:"5px"}}),precision:0,suffix:"kbps"})})})}),(0,a.jsx)(c.Z,{children:(0,a.jsx)(u.Z,{type:"inner",children:(0,a.jsx)("div",{style:Ee,children:(0,a.jsx)(d.Z,{title:"Viewer Latency",value:"".concat(Fe),prefix:(0,a.jsx)(w.Z,{style:{marginRight:"5px"}}),precision:0,suffix:"seconds"})})})}),(0,a.jsx)(c.Z,{children:(0,a.jsx)(u.Z,{type:"inner",children:(0,a.jsxs)("div",{style:Ee,children:[(0,a.jsx)(d.Z,{title:"Recent Playback Errors",value:"".concat(Ce||0),valueStyle:{color:Ne},prefix:(0,a.jsx)(x.Z,{style:{marginRight:"5px"}}),suffix:""})," "]})})})]}),(0,a.jsx)("div",{style:{textAlign:"center",display:ue<100?"grid":"none"},children:(0,a.jsxs)(o.Z.Paragraph,{style:{opacity:"0.7",fontSize:"0.7em"},children:["Provided playback metrics represent ",ue,"% of all known players. Specifics around other players consuming this video stream are unknown."]})}),(0,a.jsxs)(u.Z,{children:[(0,a.jsx)(F,{title:"Video Segment Download",description:(0,a.jsxs)(a.Fragment,{children:[(0,a.jsx)(o.Z.Paragraph,{children:"Once a video segment takes too long to download a viewer will experience buffering. If you see slow downloads you should offer a lower quality for your viewers, or find other ways, possibly an external storage provider, a CDN or a faster network, to improve your stream quality. Increasing your latency buffer can also help for some viewers."}),(0,a.jsx)(o.Z.Paragraph,{children:"In short, once the pink line consistently gets near the blue line, your stream is likely experiencing problems for viewers."})]})}),Se&&(0,a.jsx)(f.Z,{message:"Slow downloads",description:Se,type:"error",showIcon:!0}),(0,a.jsx)(j.Z,{title:"Seconds",dataCollections:ve,color:"#FF7700",unit:"s",yLogarithmic:!0})]}),(0,a.jsxs)(u.Z,{children:[(0,a.jsx)(F,{title:"Player Network Speed",description:(0,a.jsxs)(a.Fragment,{children:[(0,a.jsx)(o.Z.Paragraph,{children:"The playback bitrate of your viewers. Once somebody's bitrate drops below the lowest video variant bitrate they will experience buffering. If you see viewers with slow connections trying to play your video you should consider offering an additional, lower quality."}),(0,a.jsx)(o.Z.Paragraph,{children:"In short, once the pink line gets near the lowest blue line, your stream is likely experiencing problems for at least one of your viewers."})]})}),Pe&&(0,a.jsx)(f.Z,{message:"Low bandwidth viewers",description:Pe,type:"error",showIcon:!0}),(0,a.jsx)(j.Z,{title:"Lowest Player Bitrate",dataCollections:ye,color:"#FF7700",unit:"kbps",yLogarithmic:!0})]}),(0,a.jsxs)(u.Z,{children:[(0,a.jsx)(F,{title:"Errors and Quality Changes",description:(0,a.jsxs)(a.Fragment,{children:[(0,a.jsx)(o.Z.Paragraph,{children:"Recent number of errors, including buffering, and quality changes from across all your viewers. Errors can occur for many reasons, including browser issues, plugins, wifi problems, and they don't all represent fatal issues or something you have control over."}),"A quality change is not necessarily a negative thing, but if it's excessive and coinciding with errors you should consider adding another quality variant.",(0,a.jsx)(o.Z.Paragraph,{})]})}),(0,a.jsx)(j.Z,{title:"#",dataCollections:he,color:"#FF7700",unit:""})]}),(0,a.jsxs)(u.Z,{children:[(0,a.jsx)(F,{title:"Viewer Latency",description:"An approximate number of seconds that your viewers are behind your live video. The more people buffer the further behind they will be. High latency itself is not a problem, but some people care about this more than others."}),(0,a.jsx)(j.Z,{title:"Seconds",dataCollections:me,color:"#FF7700",unit:"s"})]})]})]})}},56180:function(e,t,n){"use strict";n.d(t,{Z:function(){return v}});var r=n(4942),i=n(1413),a=n(97685),o=n(91),s=n(67294),l=n(82532),c=n(94184),u=n.n(c),d={adjustX:1,adjustY:1},f=[0,0],p={topLeft:{points:["bl","tl"],overflow:d,offset:[0,-4],targetOffset:f},topCenter:{points:["bc","tc"],overflow:d,offset:[0,-4],targetOffset:f},topRight:{points:["br","tr"],overflow:d,offset:[0,-4],targetOffset:f},bottomLeft:{points:["tl","bl"],overflow:d,offset:[0,4],targetOffset:f},bottomCenter:{points:["tc","bc"],overflow:d,offset:[0,4],targetOffset:f},bottomRight:{points:["tr","br"],overflow:d,offset:[0,4],targetOffset:f}},h=["arrow","prefixCls","transitionName","animation","align","placement","placements","getPopupContainer","showAction","hideAction","overlayClassName","overlayStyle","visible","trigger"];function m(e,t){var n=e.arrow,c=void 0!==n&&n,d=e.prefixCls,f=void 0===d?"rc-dropdown":d,m=e.transitionName,v=e.animation,y=e.align,g=e.placement,w=void 0===g?"bottomLeft":g,x=e.placements,b=void 0===x?p:x,j=e.getPopupContainer,Z=e.showAction,F=e.hideAction,C=e.overlayClassName,k=e.overlayStyle,P=e.visible,S=e.trigger,N=void 0===S?["hover"]:S,E=(0,o.Z)(e,h),O=s.useState(),L=(0,a.Z)(O,2),D=L[0],M=L[1],T="visible"in e?P:D,A=s.useRef(null);s.useImperativeHandle(t,(function(){return A.current}));var q=function(){var t=e.overlay;return"function"===typeof t?t():t},I=function(t){var n=e.onOverlayClick,r=q().props;M(!1),n&&n(t),r.onClick&&r.onClick(t)},z=function(){var e=q(),t={prefixCls:"".concat(f,"-menu"),onClick:I};return"string"===typeof e.type&&delete t.prefixCls,s.createElement(s.Fragment,null,c&&s.createElement("div",{className:"".concat(f,"-arrow")}),s.cloneElement(e,t))},B=F;return B||-1===N.indexOf("contextMenu")||(B=["click"]),s.createElement(l.Z,(0,i.Z)((0,i.Z)({builtinPlacements:b},E),{},{prefixCls:f,ref:A,popupClassName:u()(C,(0,r.Z)({},"".concat(f,"-show-arrow"),c)),popupStyle:k,action:N,showAction:Z,hideAction:B||[],popupPlacement:w,popupAlign:y,popupTransitionName:m,popupAnimation:v,popupVisible:T,stretch:function(){var t=e.minOverlayWidthMatchTrigger,n=e.alignPoint;return"minOverlayWidthMatchTrigger"in e?t:!n}()?"minWidth":"",popup:"function"===typeof e.overlay?z:z(),onPopupVisibleChange:function(t){var n=e.onVisibleChange;M(t),"function"===typeof n&&n(t)},getPopupContainer:j}),function(){var t=e.children,n=t.props?t.props:{},r=u()(n.className,function(){var t=e.openClassName;return void 0!==t?t:"".concat(f,"-open")}());return T&&t?s.cloneElement(t,{className:r}):t}())}var v=s.forwardRef(m)}},function(e){e.O(0,[570,91,879,751,763,80,774,888,179],(function(){return t=86401,e(e.s=t);var t}));var t=e.O();_N_E=t}]);