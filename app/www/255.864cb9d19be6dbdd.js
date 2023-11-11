"use strict";(self.webpackChunkapp=self.webpackChunkapp||[]).push([[255],{2255:(E,R,L)=>{L.r(R),L.d(R,{startInputShims:()=>q});var g=L(5861),T=L(3679),m=L(3630),N=L(6991),_=L(3178);const M=new WeakMap,P=(e,t,r,s=0,o=!1)=>{M.has(e)!==r&&(r?Z(e,t,s,o):G(e,t))},Z=(e,t,r,s=!1)=>{const o=t.parentNode,n=t.cloneNode(!1);n.classList.add("cloned-input"),n.tabIndex=-1,s&&(n.disabled=!0),o.appendChild(n),M.set(e,n);const i="rtl"===e.ownerDocument.dir?9999:-9999;e.style.pointerEvents="none",t.style.transform=`translate3d(${i}px,${r}px,0) scale(0)`},G=(e,t)=>{const r=M.get(e);r&&(M.delete(e),r.remove()),e.style.pointerEvents="",t.style.transform=""},C="input, textarea, [no-blur], [contenteditable]",F="$ionPaddingTimer",B=(e,t,r)=>{const s=e[F];s&&clearTimeout(s),t>0?e.style.setProperty("--keyboard-offset",`${t}px`):e[F]=setTimeout(()=>{e.style.setProperty("--keyboard-offset","0px"),r&&r()},120)},U=(e,t,r)=>{e.addEventListener("focusout",()=>{t&&B(t,0,r)},{once:!0})};let b=0;const p="data-ionic-skip-scroll-assist",V=(e,t,r,s,o,n,a,i=!1)=>{const l=n&&(void 0===a||a.mode===N.a.None);let y=!1;const u=void 0!==_.w?_.w.innerHeight:0,v=S=>{!1!==y?H(e,t,r,s,S.detail.keyboardHeight,l,i,u,!1):y=!0},c=()=>{y=!1,null==_.w||_.w.removeEventListener("ionKeyboardDidShow",v),e.removeEventListener("focusout",c,!0)},f=function(){var S=(0,g.Z)(function*(){t.hasAttribute(p)?t.removeAttribute(p):(H(e,t,r,s,o,l,i,u),null==_.w||_.w.addEventListener("ionKeyboardDidShow",v),e.addEventListener("focusout",c,!0))});return function(){return S.apply(this,arguments)}}();return e.addEventListener("focusin",f,!0),()=>{e.removeEventListener("focusin",f,!0),null==_.w||_.w.removeEventListener("ionKeyboardDidShow",v),e.removeEventListener("focusout",c,!0)}},x=e=>{document.activeElement!==e&&(e.setAttribute(p,"true"),e.focus())},H=function(){var e=(0,g.Z)(function*(t,r,s,o,n,a,i=!1,l=0,y=!0){if(!s&&!o)return;const u=((e,t,r,s)=>{var o;return((e,t,r,s)=>{const o=e.top,n=e.bottom,a=t.top,l=a+15,u=Math.min(t.bottom,s-r)-50-n,v=l-o,c=Math.round(u<0?-u:v>0?-v:0),f=Math.min(c,o-a),w=Math.abs(f)/.3;return{scrollAmount:f,scrollDuration:Math.min(400,Math.max(150,w)),scrollPadding:r,inputSafeY:4-(o-l)}})((null!==(o=e.closest("ion-item,[ion-item]"))&&void 0!==o?o:e).getBoundingClientRect(),t.getBoundingClientRect(),r,s)})(t,s||o,n,l);if(s&&Math.abs(u.scrollAmount)<4)return x(r),void(a&&null!==s&&(B(s,b),U(r,s,()=>b=0)));if(P(t,r,!0,u.inputSafeY,i),x(r),(0,m.r)(()=>t.click()),a&&s&&(b=u.scrollPadding,B(s,b)),typeof window<"u"){let v;const c=function(){var S=(0,g.Z)(function*(){void 0!==v&&clearTimeout(v),window.removeEventListener("ionKeyboardDidShow",f),window.removeEventListener("ionKeyboardDidShow",c),s&&(yield(0,T.c)(s,0,u.scrollAmount,u.scrollDuration)),P(t,r,!1,u.inputSafeY),x(r),a&&U(r,s,()=>b=0)});return function(){return S.apply(this,arguments)}}(),f=()=>{window.removeEventListener("ionKeyboardDidShow",f),window.addEventListener("ionKeyboardDidShow",c)};if(s){const S=yield(0,T.g)(s);if(y&&u.scrollAmount>S.scrollHeight-S.clientHeight-S.scrollTop)return"password"===r.type?(u.scrollAmount+=50,window.addEventListener("ionKeyboardDidShow",f)):window.addEventListener("ionKeyboardDidShow",c),void(v=setTimeout(c,1e3))}c()}});return function(r,s,o,n,a,i){return e.apply(this,arguments)}}(),q=function(){var e=(0,g.Z)(function*(t,r){const s=document,o="ios"===r,n="android"===r,a=t.getNumber("keyboardHeight",290),i=t.getBoolean("scrollAssist",!0),l=t.getBoolean("hideCaretOnScroll",o),y=t.getBoolean("inputBlurring",o),u=t.getBoolean("scrollPadding",!0),v=Array.from(s.querySelectorAll("ion-input, ion-textarea")),c=new WeakMap,f=new WeakMap,S=yield N.K.getResizeMode(),w=function(){var h=(0,g.Z)(function*(d){yield new Promise(I=>(0,m.c)(d,I));const O=d.shadowRoot||d,D=O.querySelector("input")||O.querySelector("textarea"),A=(0,T.a)(d),W=A?null:d.closest("ion-footer");if(D){if(A&&l&&!c.has(d)){const I=((e,t,r)=>{if(!r||!t)return()=>{};const s=i=>{(e=>e===e.getRootNode().activeElement)(t)&&P(e,t,i)},o=()=>P(e,t,!1),n=()=>s(!0),a=()=>s(!1);return(0,m.a)(r,"ionScrollStart",n),(0,m.a)(r,"ionScrollEnd",a),t.addEventListener("blur",o),()=>{(0,m.b)(r,"ionScrollStart",n),(0,m.b)(r,"ionScrollEnd",a),t.removeEventListener("blur",o)}})(d,D,A);c.set(d,I)}if("date"!==D.type&&"datetime-local"!==D.type&&(A||W)&&i&&!f.has(d)){const I=V(d,D,A,W,a,u,S,n);f.set(d,I)}}});return function(O){return h.apply(this,arguments)}}();y&&(()=>{let e=!0,t=!1;const r=document;(0,m.a)(r,"ionScrollStart",()=>{t=!0}),r.addEventListener("focusin",()=>{e=!0},!0),r.addEventListener("touchend",a=>{if(t)return void(t=!1);const i=r.activeElement;if(!i||i.matches(C))return;const l=a.target;l!==i&&(l.matches(C)||l.closest(C)||(e=!1,setTimeout(()=>{e||i.blur()},50)))},!1)})();for(const h of v)w(h);s.addEventListener("ionInputDidLoad",h=>{w(h.detail)}),s.addEventListener("ionInputDidUnload",h=>{(h=>{if(l){const d=c.get(h);d&&d(),c.delete(h)}if(i){const d=f.get(h);d&&d(),f.delete(h)}})(h.detail)})});return function(r,s){return e.apply(this,arguments)}}()}}]);