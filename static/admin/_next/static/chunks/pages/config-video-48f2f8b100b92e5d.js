(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[64],{79893:function(e,n,t){(window.__NEXT_P=window.__NEXT_P||[]).push(["/config-video",function(){return t(37252)}])},15976:function(e,n,t){"use strict";t.d(n,{Z:function(){return f}});var r=t(34051),a=t.n(r),i=t(85893),o=t(67294),s=t(94594),l=t(83200),c=t(78464),u=t(25964),d=t(35159);function h(e,n,t,r,a,i,o){try{var s=e[i](o),l=s.value}catch(c){return void t(c)}s.done?n(l):Promise.resolve(l).then(r,a)}function f(e){var n=(0,o.useState)(null),t=n[0],r=n[1],f=null,m=((0,o.useContext)(d.aC)||{}).setFieldInConfigState,v=e.apiPath,p=e.checked,g=e.reversed,y=void 0!==g&&g,x=e.configPath,b=void 0===x?"":x,j=e.disabled,w=void 0!==j&&j,N=e.fieldName,k=e.label,P=e.tip,C=e.useSubmit,S=e.onChange,Z=function(){r(null),clearTimeout(f),f=null},I=function(){var e,n=(e=a().mark((function e(n){var t;return a().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:if(!C){e.next=6;break}return r((0,l.kg)(l.Jk)),t=y?!n:n,e.next=5,(0,u.Si)({apiPath:v,data:{value:t},onSuccess:function(){m({fieldName:N,value:t,path:b}),r((0,l.kg)(l.zv))},onError:function(e){r((0,l.kg)(l.Un,"There was an error: ".concat(e)))}});case 5:f=setTimeout(Z,u.sI);case 6:S&&S(n);case 7:case"end":return e.stop()}}),e)})),function(){var n=this,t=arguments;return new Promise((function(r,a){var i=e.apply(n,t);function o(e){h(i,r,a,o,s,"next",e)}function s(e){h(i,r,a,o,s,"throw",e)}o(void 0)}))});return function(e){return n.apply(this,arguments)}}(),T=null!==t&&t.type===l.Jk;return(0,i.jsxs)("div",{className:"formfield-container toggleswitch-container",children:[k&&(0,i.jsx)("div",{className:"label-side",children:(0,i.jsx)("span",{className:"formfield-label",children:k})}),(0,i.jsxs)("div",{className:"input-side",children:[(0,i.jsxs)("div",{className:"input-group",children:[(0,i.jsx)(s.Z,{className:"switch field-".concat(N),loading:T,onChange:I,defaultChecked:p,checked:p,checkedChildren:"ON",unCheckedChildren:"OFF",disabled:w}),(0,i.jsx)(c.Z,{status:t})]}),(0,i.jsx)("p",{className:"field-tip",children:P})]})]})}f.defaultProps={apiPath:"",checked:!1,reversed:!1,configPath:"",disabled:!1,label:"",tip:"",useSubmit:!1,onChange:null}},37252:function(e,n,t){"use strict";t.r(n),t.d(n,{default:function(){return ee}});var r=t(85893),a=t(54907),i=t(84916),o=t(25968),s=t(6226),l=t(67294),c=t(34051),u=t.n(c),d=t(38939),h=t(75443),f=t(57553),m=t(25964),v=t(83200),p=t(35159),g=t(78464);function y(e,n){(null==n||n>e.length)&&(n=e.length);for(var t=0,r=new Array(n);t<n;t++)r[t]=e[t];return r}function x(e,n,t,r,a,i,o){try{var s=e[i](o),l=s.value}catch(c){return void t(c)}s.done?n(l):Promise.resolve(l).then(r,a)}function b(e){return function(){var n=this,t=arguments;return new Promise((function(r,a){var i=e.apply(n,t);function o(e){x(i,r,a,o,s,"next",e)}function s(e){x(i,r,a,o,s,"throw",e)}o(void 0)}))}}function j(e,n){return function(e){if(Array.isArray(e))return e}(e)||function(e,n){var t=null==e?null:"undefined"!==typeof Symbol&&e[Symbol.iterator]||e["@@iterator"];if(null!=t){var r,a,i=[],o=!0,s=!1;try{for(t=t.call(e);!(o=(r=t.next()).done)&&(i.push(r.value),!n||i.length!==n);o=!0);}catch(l){s=!0,a=l}finally{try{o||null==t.return||t.return()}finally{if(s)throw a}}return i}}(e,n)||function(e,n){if(!e)return;if("string"===typeof e)return y(e,n);var t=Object.prototype.toString.call(e).slice(8,-1);"Object"===t&&e.constructor&&(t=e.constructor.name);if("Map"===t||"Set"===t)return Array.from(t);if("Arguments"===t||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(t))return y(e,n)}(e,n)||function(){throw new TypeError("Invalid attempt to destructure non-iterable instance.\\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}()}function w(){var e=(0,l.useContext)(p.aC),n=e||{},t=n.serverConfig,a=n.setFieldInConfigState,o=t||{},s=o.videoCodec,c=o.supportedCodecs,y=i.Z.Title,x=d.Z.Option,w=(0,l.useState)(null),N=w[0],k=w[1],P=(0,l.useContext)(f.k).setMessage,C=(0,l.useState)(s),S=C[0],Z=C[1],I=(0,l.useState)(s),T=I[0],_=I[1],V=j(l.useState(!1),2),O=V[0],A=V[1],U=null;(0,l.useEffect)((function(){Z(s)}),[s]);var E=function(){k(null),U=null,clearTimeout(U)};function L(){return(L=b(u().mark((function n(){return u().wrap((function(n){for(;;)switch(n.prev=n.next){case 0:return Z(T),_(""),A(!1),n.next=5,(0,m.Si)({apiPath:m.CQ,data:{value:T},onSuccess:function(){a({fieldName:"videoCodec",value:T,path:"videoSettings"}),k((0,v.kg)(v.zv,"Video codec updated.")),U=setTimeout(E,m.sI),e.online&&P("Your latency buffer setting will take effect the next time you begin a live stream.")},onError:function(e){k((0,v.kg)(v.Un,e)),U=setTimeout(E,m.sI)}});case 5:case"end":return n.stop()}}),n)})))).apply(this,arguments)}var B=c.map((function(e){var n=e;return"libx264"===n?n="Default (libx264)":"h264_nvenc"===n?n="NVIDIA GPU acceleration":"h264_vaapi"===n?n="VA-API hardware encoding":"h264_qsv"===n?n="Intel QuickSync":"h264_v4l2m2m"===n?n="Video4Linux hardware encoding":"h264_omx"===n&&(n="OpenMax (omx) for Raspberry Pi"),(0,r.jsx)(x,{value:e,children:n},e)})),F="";return"libx264"===S?F="libx264 is the default codec and generally the only working choice for shared VPS enviornments. This is likely what you should be using unless you know you have set up other options.":"h264_nvenc"===S?F="You can use your NVIDIA GPU for encoding if you have a modern NVIDIA card with encoding cores.":"h264_vaapi"===S?F="VA-API may be supported by your NVIDIA proprietary drivers, Mesa open-source drivers for AMD or Intel graphics.":"h264_qsv"===S?F="Quick Sync Video is Intel's brand for its dedicated video encoding and decoding hardware. It may be an option if you have a modern Intel CPU with integrated graphics.":"h264_v4l2m2m"===S?F="Video4Linux is an interface to multiple different hardware encoding platforms such as Intel and AMD.":"h264_omx"===S&&(F="OpenMax is a codec most often used with a Raspberry Pi."),(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(y,{level:3,className:"section-title",children:"Video Codec"}),(0,r.jsxs)("div",{className:"description",children:["If you have access to specific hardware with the drivers and software installed for them, you may be able to improve your video encoding performance.",(0,r.jsx)("p",{children:(0,r.jsx)("a",{href:"https://owncast.online/docs/codecs?source=admin",target:"_blank",rel:"noopener noreferrer",children:"Read the documentation about this setting before changing it or you may make your stream unplayable."})})]}),(0,r.jsxs)("div",{className:"segment-slider-container",children:[(0,r.jsx)(h.Z,{title:"Are you sure you want to change your video codec to ".concat(T," and understand what this means?"),visible:O,placement:"leftBottom",onConfirm:function(){return L.apply(this,arguments)},okText:"Yes",cancelText:"No",children:(0,r.jsx)(d.Z,{defaultValue:S,value:S,style:{width:"100%"},onChange:function(e){_(e),A(!0)},children:B})}),(0,r.jsx)(g.Z,{status:N}),(0,r.jsx)("p",{id:"selected-codec-note",className:"selected-value-note",children:F})]})]})}var N=t(23028);function k(e,n,t,r,a,i,o){try{var s=e[i](o),l=s.value}catch(c){return void t(c)}s.done?n(l):Promise.resolve(l).then(r,a)}var P=i.Z.Title,C={0:"Lowest",1:"",2:"",3:"",4:"Highest"},S={0:"Lowest latency, lowest error tolerance (Not recommended, may not work for all content/configurations.)",1:"Low latency, low error tolerance",2:"Medium latency, medium error tolerance (Default)",3:"High latency, high error tolerance",4:"Highest latency, highest error tolerance"};function Z(){var e=(0,l.useState)(null),n=e[0],t=e[1],a=(0,l.useState)(null),i=a[0],o=a[1],s=(0,l.useContext)(p.aC),c=(0,l.useContext)(f.k).setMessage,d=s||{},h=d.serverConfig,y=d.setFieldInConfigState,x=(h||{}).videoSettings,b=null;if(!x)return null;(0,l.useEffect)((function(){o(x.latencyLevel)}),[x]);var j=function(){t(null),b=null,clearTimeout(b)},w=function(){var e,n=(e=u().mark((function e(n){return u().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return t((0,v.kg)(v.Jk)),e.next=3,(0,m.Si)({apiPath:m.sv,data:{value:n},onSuccess:function(){y({fieldName:"latencyLevel",value:n,path:"videoSettings"}),t((0,v.kg)(v.zv,"Latency buffer level updated.")),b=setTimeout(j,m.sI),s.online&&c("Your latency buffer setting will take effect the next time you begin a live stream.")},onError:function(e){t((0,v.kg)(v.Un,e)),b=setTimeout(j,m.sI)}});case 3:case"end":return e.stop()}}),e)})),function(){var n=this,t=arguments;return new Promise((function(r,a){var i=e.apply(n,t);function o(e){k(i,r,a,o,s,"next",e)}function s(e){k(i,r,a,o,s,"throw",e)}o(void 0)}))});return function(e){return n.apply(this,arguments)}}();return(0,r.jsxs)("div",{className:"config-video-latency-container",children:[(0,r.jsx)(P,{level:3,className:"section-title",children:"Latency Buffer"}),(0,r.jsx)("p",{className:"description",children:"While it's natural to want to keep your latency as low as possible, you may experience reduced error tolerance and stability the lower you go. The lowest setting is not recommended."}),(0,r.jsxs)("p",{className:"description",children:["For interactive live streams you may want to experiment with a lower latency, for non-interactive broadcasts you may want to increase it."," ",(0,r.jsx)("a",{href:"https://owncast.online/docs/encoding#latency-buffer?source=admin",target:"_blank",rel:"noopener noreferrer",children:"Read to learn more."})]}),(0,r.jsxs)("div",{className:"segment-slider-container",children:[(0,r.jsx)(N.Z,{tipFormatter:function(e){return S[e]},onChange:function(e){w(e)},min:0,max:4,marks:C,defaultValue:i,value:i}),(0,r.jsx)("p",{className:"selected-value-note",children:S[i]}),(0,r.jsx)(g.Z,{status:n})]})]})}var I=t(71577),T=t(14670),_=t(6551),V=t(52455),O=t(48689),A=t(21640),U=t(94184),E=t.n(U),L=t(48419),B=t(15976);function F(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function D(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{},r=Object.keys(t);"function"===typeof Object.getOwnPropertySymbols&&(r=r.concat(Object.getOwnPropertySymbols(t).filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable})))),r.forEach((function(n){F(e,n,t[n])}))}return e}var R=a.Z.Panel;function M(e){var n=e.dataState,t=void 0===n?m.gX:n,l=e.onUpdateField,c=t.videoPassthrough,u=E()({"config-variant-form":!0,"video-passthrough-enabled":c});return(0,r.jsxs)("div",{className:u,children:[(0,r.jsxs)("p",{className:"description",children:[(0,r.jsx)("a",{href:"https://owncast.online/docs/video?source=admin",target:"_blank",rel:"noopener noreferrer",children:"Learn more"})," ","about how each of these settings can impact the performance of your server."]}),c&&(0,r.jsxs)("p",{className:"passthrough-warning",children:["NOTE: Video Passthrough for this output stream variant is ",(0,r.jsx)("em",{children:"enabled"}),", disabling the below video encoding settings."]}),(0,r.jsxs)(o.Z,{gutter:16,children:[(0,r.jsx)(L.ZP,D({maxLength:"10"},m.SS,{value:t.name,onChange:function(e){l({fieldName:"name",value:e.value})}})),(0,r.jsx)(s.Z,{sm:24,md:12,children:(0,r.jsxs)("div",{className:"form-module cpu-usage-container",children:[(0,r.jsx)(i.Z.Title,{level:3,children:"CPU or GPU Utilization"}),(0,r.jsx)("p",{className:"description",children:"Reduce to improve server performance, or increase it to improve video quality."}),(0,r.jsxs)("div",{className:"segment-slider-container",children:[(0,r.jsx)(N.Z,{tipFormatter:function(e){return m.I$[e]},onChange:function(e){l({fieldName:"cpuUsageLevel",value:e})},min:1,max:Object.keys(m.t$).length,marks:m.t$,defaultValue:t.cpuUsageLevel,value:t.cpuUsageLevel,disabled:t.videoPassthrough}),(0,r.jsx)("p",{className:"selected-value-note",children:c?"CPU usage selection is disabled when Video Passthrough is enabled.":m.I$[t.cpuUsageLevel]||""})]}),(0,r.jsxs)("p",{className:"read-more-subtext",children:["This could mean GPU or CPU usage depending on your server environment."," ",(0,r.jsx)("a",{href:"https://owncast.online/docs/video/?source=admin#cpu-usage",target:"_blank",rel:"noopener noreferrer",children:"Read more about hardware performance."})]})]})}),(0,r.jsx)(s.Z,{sm:24,md:12,children:(0,r.jsxs)("div",{className:"form-module bitrate-container ".concat(t.videoPassthrough?"disabled":""),children:[(0,r.jsx)(i.Z.Title,{level:3,children:"Video Bitrate"}),(0,r.jsx)("p",{className:"description",children:m.yC.tip}),(0,r.jsxs)("div",{className:"segment-slider-container",children:[(0,r.jsx)(N.Z,{tipFormatter:function(e){return"".concat(e," ").concat(m.yC.unit)},disabled:t.videoPassthrough,defaultValue:t.videoBitrate,value:t.videoBitrate,onChange:function(e){l({fieldName:"videoBitrate",value:e})},step:m.yC.incrementBy,min:m.yC.min,max:m.yC.max,marks:m.HM}),(0,r.jsx)("p",{className:"selected-value-note",children:function(){if(c)return"Bitrate selection is disabled when Video Passthrough is enabled.";var e="".concat(t.videoBitrate).concat(m.yC.unit);return e=t.videoBitrate<2e3?"".concat(e," - Good for low bandwidth environments."):t.videoBitrate<3500?"".concat(e," - Good for most bandwidth environments."):"".concat(e," - Good for high bandwidth environments.")}()})]}),(0,r.jsx)("p",{className:"read-more-subtext",children:(0,r.jsx)("a",{href:"https://owncast.online/docs/video/?source=admin",target:"_blank",rel:"noopener noreferrer",children:"Read more about bitrates."})})]})})]}),(0,r.jsx)(a.Z,{className:"advanced-settings",children:(0,r.jsxs)(R,{header:"Advanced Settings",children:[(0,r.jsxs)(o.Z,{gutter:16,children:[(0,r.jsx)(s.Z,{sm:24,md:12,children:(0,r.jsxs)("div",{className:"form-module resolution-module",children:[(0,r.jsx)(i.Z.Title,{level:3,children:"Resolution"}),(0,r.jsxs)("p",{className:"description",children:["Resizing your content will take additional resources on your server. If you wish to optionally resize your content for this stream output then you should either set the width ",(0,r.jsx)("strong",{children:"or"})," the height to keep your aspect ratio."," ",(0,r.jsx)("a",{href:"https://owncast.online/docs/video/?source=admin",target:"_blank",rel:"noopener noreferrer",children:"Read more about resolutions."})]}),(0,r.jsx)("br",{}),(0,r.jsx)(L.ZP,D({type:"number"},m.dL.scaledWidth,{value:t.scaledWidth,onChange:function(e){var n=Number(e.value);isNaN(n)||l({fieldName:"scaledWidth",value:n||0})},disabled:t.videoPassthrough})),(0,r.jsx)(L.ZP,D({type:"number"},m.dL.scaledHeight,{value:t.scaledHeight,onChange:function(e){var n=Number(e.value);isNaN(n)||l({fieldName:"scaledHeight",value:n||0})},disabled:t.videoPassthrough}))]})}),(0,r.jsx)(s.Z,{sm:24,md:12,children:(0,r.jsxs)("div",{className:"form-module video-passthrough-module",children:[(0,r.jsx)(i.Z.Title,{level:3,children:"Video Passthrough"}),(0,r.jsxs)("div",{className:"description",children:[(0,r.jsxs)("p",{children:["Enabling video passthrough may allow for less hardware utilization, but may also make your stream ",(0,r.jsx)("strong",{children:"unplayable"}),"."]}),(0,r.jsx)("p",{children:"All other settings for this stream output will be disabled if passthrough is used."}),(0,r.jsx)("p",{children:(0,r.jsx)("a",{href:"https://owncast.online/docs/video/?source=admin#video-passthrough",target:"_blank",rel:"noopener noreferrer",children:"Read the documentation before enabling, as it impacts your stream."})})]}),(0,r.jsx)(h.Z,{disabled:!0===t.videoPassthrough,title:"Did you read the documentation about video passthrough and understand the risks involved with enabling it?",icon:(0,r.jsx)(A.Z,{}),onConfirm:function(){l({fieldName:"videoPassthrough",value:!0})},okText:"Yes",cancelText:"No",children:(0,r.jsx)("a",{href:"#",children:(0,r.jsx)(B.Z,{label:"Use Video Passthrough?",fieldName:"video-passthrough",tip:m.dL.videoPassthrough.tip,checked:t.videoPassthrough,onChange:function(e){c&&l({fieldName:"videoPassthrough",value:e})}})})})]})})]}),(0,r.jsxs)("div",{className:"form-module frame-rate-module",children:[(0,r.jsx)(i.Z.Title,{level:3,children:"Frame rate"}),(0,r.jsx)("p",{className:"description",children:m.nm.tip}),(0,r.jsxs)("div",{className:"segment-slider-container",children:[(0,r.jsx)(N.Z,{tipFormatter:function(e){return"".concat(e," ").concat(m.nm.unit)},defaultValue:t.framerate,value:t.framerate,onChange:function(e){l({fieldName:"framerate",value:e})},step:m.nm.incrementBy,min:m.nm.min,max:m.nm.max,marks:m.Xq,disabled:t.videoPassthrough}),(0,r.jsx)("p",{className:"selected-value-note",children:c?"Framerate selection is disabled when Video Passthrough is enabled.":m.x8[t.framerate]||""})]}),(0,r.jsx)("p",{className:"read-more-subtext",children:(0,r.jsx)("a",{href:"https://owncast.online/docs/video/?source=admin#framerate",target:"_blank",rel:"noopener noreferrer",children:"Read more about framerates."})})]})]},"1")})]})}function z(e,n){(null==n||n>e.length)&&(n=e.length);for(var t=0,r=new Array(n);t<n;t++)r[t]=e[t];return r}function H(e,n,t,r,a,i,o){try{var s=e[i](o),l=s.value}catch(c){return void t(c)}s.done?n(l):Promise.resolve(l).then(r,a)}function G(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function X(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{},r=Object.keys(t);"function"===typeof Object.getOwnPropertySymbols&&(r=r.concat(Object.getOwnPropertySymbols(t).filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable})))),r.forEach((function(n){G(e,n,t[n])}))}return e}function $(e,n){return function(e){if(Array.isArray(e))return e}(e)||function(e,n){var t=null==e?null:"undefined"!==typeof Symbol&&e[Symbol.iterator]||e["@@iterator"];if(null!=t){var r,a,i=[],o=!0,s=!1;try{for(t=t.call(e);!(o=(r=t.next()).done)&&(i.push(r.value),!n||i.length!==n);o=!0);}catch(l){s=!0,a=l}finally{try{o||null==t.return||t.return()}finally{if(s)throw a}}return i}}(e,n)||Q(e,n)||function(){throw new TypeError("Invalid attempt to destructure non-iterable instance.\\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}()}function q(e){return function(e){if(Array.isArray(e))return z(e)}(e)||function(e){if("undefined"!==typeof Symbol&&null!=e[Symbol.iterator]||null!=e["@@iterator"])return Array.from(e)}(e)||Q(e)||function(){throw new TypeError("Invalid attempt to spread non-iterable instance.\\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}()}function Q(e,n){if(e){if("string"===typeof e)return z(e,n);var t=Object.prototype.toString.call(e).slice(8,-1);return"Object"===t&&e.constructor&&(t=e.constructor.name),"Map"===t||"Set"===t?Array.from(t):"Arguments"===t||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(t)?z(e,n):void 0}}var Y=i.Z.Title;function J(){var e=(0,l.useState)(!1),n=e[0],t=e[1],a=(0,l.useState)(!1),i=a[0],o=a[1],s=(0,l.useState)(0),c=s[0],d=s[1],h=(0,l.useContext)(f.k).setMessage,y=(0,l.useState)(m.gX),x=y[0],b=y[1],j=(0,l.useState)(null),w=j[0],N=j[1],k=(0,l.useContext)(p.aC),P=k||{},C=P.serverConfig,S=P.setFieldInConfigState,Z=(C||{}).videoSettings,A=(Z||{}).videoQualityVariants,U=null;if(!Z)return null;var E=function(){N(null),U=null,clearTimeout(U)},L=function(){t(!1),d(-1),b(m.gX)},B=function(){var e,n=(e=u().mark((function e(n){return u().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return N((0,v.kg)(v.Jk)),e.next=3,(0,m.Si)({apiPath:m.vv,data:{value:n},onSuccess:function(){S({fieldName:"videoQualityVariants",value:n,path:"videoSettings"}),o(!1),L(),N((0,v.kg)(v.zv,"Variants updated")),U=setTimeout(E,m.sI),k.online&&h("Updating your video configuration will take effect the next time you begin a new stream.")},onError:function(e){N((0,v.kg)(v.Un,e)),o(!1),U=setTimeout(E,m.sI)}});case 3:case"end":return e.stop()}}),e)})),function(){var n=this,t=arguments;return new Promise((function(r,a){var i=e.apply(n,t);function o(e){H(i,r,a,o,s,"next",e)}function s(e){H(i,r,a,o,s,"throw",e)}o(void 0)}))});return function(e){return n.apply(this,arguments)}}(),F=[{title:"Name",dataIndex:"name",render:function(e){return e||"No name"}},{title:"Video bitrate",dataIndex:"videoBitrate",key:"videoBitrate",render:function(e,n){return!e||n.videoPassthrough?"Same as source":"".concat(e," kbps")}},{title:"CPU Usage",dataIndex:"cpuUsageLevel",key:"cpuUsageLevel",render:function(e,n){return!e||n.videoPassthrough?"n/a":m.I$[e].split(" ")[0]}},{title:"",dataIndex:"",key:"edit",render:function(e){var n=e.key-1;return(0,r.jsxs)("span",{className:"actions",children:[(0,r.jsx)(I.Z,{size:"small",onClick:function(){d(n),b(A[n]),t(!0)},children:"Edit"}),(0,r.jsx)(I.Z,{className:"delete-button",icon:(0,r.jsx)(O.Z,{}),size:"small",disabled:1===A.length,onClick:function(){!function(e){var n=q(A);n.splice(e,1),B(n)}(n)}})]})}}],D=A.map((function(e,n){return X({key:n+1},e)}));return(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(Y,{level:3,className:"section-title",children:"Stream output"}),function(){if(1!==A.length)return!1;var e=$(A,1)[0];return m.i3.VIDEO_HEIGHT<=e.scaledHeight||m.i3.VIDEO_BITRATE<=e.videoBitrate}()&&(0,r.jsx)(T.Z,{message:m.i3.HELP_TEXT,type:"info",closable:!0}),(0,r.jsx)(g.Z,{status:w}),(0,r.jsx)(_.Z,{className:"variants-table",pagination:!1,size:"small",columns:F,dataSource:D}),(0,r.jsxs)(V.Z,{title:"Edit Video Variant Details",visible:n,onOk:function(){o(!0);var e=q(A);-1===c?e.push(x):e.splice(c,1,x),B(e)},onCancel:L,confirmLoading:i,width:900,children:[(0,r.jsx)(M,{dataState:X({},x),onUpdateField:function(e){var n=e.fieldName,t=e.value;b(X({},x,G({},n,t)))}}),(0,r.jsx)(g.Z,{status:w})]}),(0,r.jsx)("br",{}),(0,r.jsx)(I.Z,{type:"primary",onClick:function(){d(-1),b(m.gX),t(!0)},children:"Add a new variant"})]})}var W=a.Z.Panel,K=i.Z.Title;function ee(){return(0,r.jsxs)("div",{className:"config-video-variants",children:[(0,r.jsx)(K,{children:"Video configuration"}),(0,r.jsxs)("p",{className:"description",children:["Before changing your video configuration"," ",(0,r.jsx)("a",{href:"https://owncast.online/docs/video?source=admin",target:"_blank",rel:"noopener noreferrer",children:"visit the video documentation"})," ","to learn how it impacts your stream performance. The general rule is to start conservatively by having one middle quality stream output variant and experiment with adding more of varied qualities."]}),(0,r.jsxs)(o.Z,{gutter:[16,16],children:[(0,r.jsx)(s.Z,{md:24,lg:12,children:(0,r.jsx)("div",{className:"form-module variants-table-module",children:(0,r.jsx)(J,{})})}),(0,r.jsxs)(s.Z,{md:24,lg:12,children:[(0,r.jsx)("div",{className:"form-module latency-module",children:(0,r.jsx)(Z,{})}),(0,r.jsx)(a.Z,{className:"advanced-settings codec-module",children:(0,r.jsx)(W,{header:"Advanced Settings",children:(0,r.jsx)("div",{className:"form-module variants-table-module",children:(0,r.jsx)(w,{})})},"1")})]})]})]})}}},function(e){e.O(0,[551,578,598,774,888,179],(function(){return n=79893,e(e.s=n);var n}));var n=e.O();_N_E=n}]);