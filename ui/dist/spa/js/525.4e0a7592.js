"use strict";(self["webpackChunkgame_stats_vue"]=self["webpackChunkgame_stats_vue"]||[]).push([[525],{4525:(e,a,t)=>{t.r(a),t.d(a,{default:()=>U});var l=t(3673),i=t(2323);const n={key:0},o={key:1,class:"column q-pa-md q-gutter-md"},s={class:"text-subtitle2"},r={class:"text-subtitle2"},d={class:"text-subtitle2 text-center"};function u(e,a,t,u,c,p){const m=(0,l.up)("q-tab"),v=(0,l.up)("q-tabs"),g=(0,l.up)("q-card-section"),f=(0,l.up)("apexchart"),h=(0,l.up)("q-card"),w=(0,l.up)("q-item-section"),y=(0,l.up)("q-item"),b=(0,l.up)("q-btn"),_=(0,l.up)("q-page-sticky"),W=(0,l.up)("q-page");return(0,l.wg)(),(0,l.j4)(W,{class:"row items-center justify-evenly"},{default:(0,l.w5)((()=>[e.game&&e.nbGames>0?((0,l.wg)(),(0,l.iD)("div",n,[(0,l.Wm)(h,{flat:"",class:"q-px-none",style:(0,i.j5)(e.$q.screen.xs?"width: 320px; height: 320px":"width: 640px; height: 460px")},{default:(0,l.w5)((()=>[(0,l.Wm)(g,{class:"q-pa-sm"},{default:(0,l.w5)((()=>[(0,l.Wm)(v,{align:"left",modelValue:e.chart,"onUpdate:modelValue":a[0]||(a[0]=a=>e.chart=a),"text-color":"dark","indicator-color":"secondary","inline-label":"",dense:""},{default:(0,l.w5)((()=>[(0,l.Wm)(m,{name:"pie",label:"Pie Chart",icon:"pie_chart"}),(0,l.Wm)(m,{name:"line",label:"Line Chart",icon:"show_chart"})])),_:1},8,["modelValue"])])),_:1}),(0,l.Wm)(g,{class:"q-px-none justify-center"},{default:(0,l.w5)((()=>["pie"==e.chart?((0,l.wg)(),(0,l.j4)(f,{key:0,width:"500",align:"center",options:e.gamePlayers,series:e.gameSeries},null,8,["options","series"])):((0,l.wg)(),(0,l.j4)(f,{key:1,width:"500",height:"385",align:"center",options:e.lineChartOptions,series:e.gamePlayersCumStats},null,8,["options","series"]))])),_:1})])),_:1},8,["style"])])):(0,l.kq)("",!0),e.game?((0,l.wg)(),(0,l.iD)("div",o,[(0,l.Wm)(h,{flat:"",bordered:"",class:"game-card text-dark",style:(0,i.j5)(e.$q.screen.xs?"width: 320px;":"width: 640px;")},{default:(0,l.w5)((()=>[(0,l.Wm)(g,{class:"q-py-none"},{default:(0,l.w5)((()=>[(0,l.Wm)(y,{class:"q-py-none"},{default:(0,l.w5)((()=>[(0,l.Wm)(w,null,{default:(0,l.w5)((()=>[(0,l._)("div",s,(0,i.zw)(e.game.name),1)])),_:1}),(0,l.Wm)(w,{avatar:""},{default:(0,l.w5)((()=>[(0,l._)("div",{class:"text-subtitle2 text-center q-px-md",style:(0,i.j5)(e.$q.screen.xs?"width: 43px;":"width: 109px;")},(0,i.zw)(e.nbGames),5)])),_:1})])),_:1})])),_:1})])),_:1},8,["style"]),((0,l.wg)(!0),(0,l.iD)(l.HY,null,(0,l.Ko)(e.game.players,(({id:a,name:t})=>((0,l.wg)(),(0,l.iD)(l.HY,{key:a},[a?((0,l.wg)(),(0,l.j4)(h,{key:0,class:"text-white bg-secondary",style:(0,i.j5)(e.$q.screen.xs?"width: 320px;":"width: 640px;")},{default:(0,l.w5)((()=>[(0,l.Wm)(g,{class:"q-py-none"},{default:(0,l.w5)((()=>[(0,l.Wm)(y,{class:"q-py-none"},{default:(0,l.w5)((()=>[(0,l.Wm)(w,null,{default:(0,l.w5)((()=>[(0,l._)("div",r,(0,i.zw)(t),1)])),_:2},1024),(0,l.Wm)(w,null,{default:(0,l.w5)((()=>[(0,l._)("div",d,(0,i.zw)(e.gamePlayersWins[t]),1)])),_:2},1024),(0,l.Wm)(w),(0,l.Wm)(w,null,{default:(0,l.w5)((()=>[(0,l.Wm)(b,{class:"full-height",flat:"","square-icon":"",icon:"add",onClick:l=>e.openAddWinnerDialog(a,t)},null,8,["onClick"])])),_:2},1024),(0,l.Wm)(w,null,{default:(0,l.w5)((()=>[(0,l.Wm)(b,{class:"full-height",size:"sm",flat:"","square-icon":"",icon:"delete",onClick:l=>e.deleteDialog(a,t)},null,8,["onClick"])])),_:2},1024)])),_:2},1024)])),_:2},1024)])),_:2},1032,["style"])):(0,l.kq)("",!0)],64)))),128))])):(0,l.kq)("",!0),(0,l.Wm)(_,{position:"bottom-right",offset:[24,24]},{default:(0,l.w5)((()=>[(0,l.Wm)(b,{fab:"",icon:"add",color:"accent",onClick:e.addPlayerDialog},null,8,["onClick"])])),_:1})])),_:1})}var c=t(1959),p=t(1768),m=t(8825);const v={class:"text-h6"},g={class:"row items-center justify-end"},f={class:"row items-center justify-end"};function h(e,a,t,n,o,s){const r=(0,l.up)("q-card-section"),d=(0,l.up)("q-btn"),u=(0,l.up)("q-date"),c=(0,l.up)("q-popup-proxy"),p=(0,l.up)("q-icon"),m=(0,l.up)("q-time"),h=(0,l.up)("q-input"),w=(0,l.up)("q-card-actions"),y=(0,l.up)("q-card"),b=(0,l.up)("q-dialog"),_=(0,l.Q2)("close-popup");return(0,l.wg)(),(0,l.j4)(b,{ref:"dialogRef",onHide:e.onDialogHide},{default:(0,l.w5)((()=>[(0,l.Wm)(y,{class:"q-dialog-plugin"},{default:(0,l.w5)((()=>[(0,l.Wm)(r,null,{default:(0,l.w5)((()=>[(0,l._)("div",v,(0,i.zw)(e.title),1)])),_:1}),(0,l.Wm)(r,{class:"q-pt-none"},{default:(0,l.w5)((()=>[(0,l.Wm)(h,{label:"Select a date",modelValue:e.datetime,"onUpdate:modelValue":a[2]||(a[2]=a=>e.datetime=a),mask:"####-##-## ##:##"},{prepend:(0,l.w5)((()=>[(0,l.Wm)(p,{name:"event",class:"cursor-pointer"},{default:(0,l.w5)((()=>[(0,l.Wm)(c,{cover:"","transition-show":"scale","transition-hide":"scale"},{default:(0,l.w5)((()=>[(0,l.Wm)(u,{modelValue:e.datetime,"onUpdate:modelValue":a[0]||(a[0]=a=>e.datetime=a),mask:"YYYY-MM-DD HH:mm"},{default:(0,l.w5)((()=>[(0,l._)("div",g,[(0,l.wy)((0,l.Wm)(d,{label:"Close",color:"primary",flat:""},null,512),[[_]])])])),_:1},8,["modelValue"])])),_:1})])),_:1})])),append:(0,l.w5)((()=>[(0,l.Wm)(p,{name:"access_time",class:"cursor-pointer"},{default:(0,l.w5)((()=>[(0,l.Wm)(c,{cover:"","transition-show":"scale","transition-hide":"scale"},{default:(0,l.w5)((()=>[(0,l.Wm)(m,{modelValue:e.datetime,"onUpdate:modelValue":a[1]||(a[1]=a=>e.datetime=a),mask:"YYYY-MM-DD HH:mm",format24h:""},{default:(0,l.w5)((()=>[(0,l._)("div",f,[(0,l.wy)((0,l.Wm)(d,{label:"Close",color:"primary",flat:""},null,512),[[_]])])])),_:1},8,["modelValue"])])),_:1})])),_:1})])),_:1},8,["modelValue"])])),_:1}),(0,l.Wm)(w,{align:"right",class:"text-primary"},{default:(0,l.w5)((()=>[(0,l.wy)((0,l.Wm)(d,{flat:"",label:"Cancel"},null,512),[[_]]),(0,l.Wm)(d,{type:"submit",flat:"",label:"Submit",onClick:e.onOKClick},null,8,["onClick"])])),_:1})])),_:1})])),_:1},8,["onHide"])}var w=t(8777),y=t(2661);const b=(0,l.aZ)({name:"AddWinnerDialog",props:{player:{type:String,required:!0},id:Number},emits:[...w.Z.emits],setup(e){const{dialogRef:a,onDialogHide:t,onDialogOK:l,onDialogCancel:i}=(0,w.Z)(),n=`Add a game win to ${e.player}`,o=Date.now(),s=(0,c.iH)(y.ZP.formatDate(o,"YYYY-MM-DD HH:mm"));return{title:n,dialogRef:a,onDialogHide:t,datetime:s,onOKClick(){const e=y.ZP.extractDate(s.value,"YYYY-MM-DD HH:mm");l(e)},onCancelClick:i}}});var _=t(4260),W=t(5926),q=t(151),k=t(5589),x=t(1206),D=t(4554),C=t(8187),Z=t(2651),P=t(8240),Q=t(1534),j=t(9367),H=t(677),Y=t(7518),O=t.n(Y);const V=(0,_.Z)(b,[["render",h]]),$=V;O()(b,"components",{QDialog:W.Z,QCard:q.Z,QCardSection:k.Z,QInput:x.Z,QIcon:D.Z,QPopupProxy:C.Z,QDate:Z.Z,QBtn:P.Z,QTime:Q.Z,QCardActions:j.Z}),O()(b,"directives",{ClosePopup:H.Z});var S=t(8603),F=t.n(S);const M=(0,l.aZ)({name:"PageGame",setup(){const e=(0,m.Z)(),a=(0,c.iH)();let t={};const l=(0,c.Fl)((()=>{var e;let t=[];return null===(e=a.value)||void 0===e||e.players.forEach((e=>{const a=e.stats.filter((e=>null!==e));t=F().union(t,a.map((e=>new Date(e).getTime())))})),t=t.sort(((e,a)=>e-a)),t})),i=(0,c.Fl)((()=>{var e;let t=[];return(null===(e=a.value)||void 0===e?void 0:e.id)&&a.value.players.forEach((e=>{const a=Object.values(Object.assign({},e.stats));let i=a.map((e=>new Date(e).getTime())).sort(((e,a)=>e-a)),n={name:e.name,data:[]},o=0;l.value.forEach((e=>{i.length>0&&i[0]===e&&(o++,i.shift()),n.data.push([e,o])})),t.push(n)})),t})),n=(0,c.Fl)((()=>{var e;let t={};return(null===(e=a.value)||void 0===e?void 0:e.id)&&a.value.players.forEach((e=>{let a=0;e.stats.forEach((e=>{e&&a++})),t[e.name]=a})),t})),o=(0,c.Fl)((()=>{var e;let t=[];return(null===(e=a.value)||void 0===e?void 0:e.id)&&a.value.players.forEach((e=>{t.push(e.name)})),Object.assign(Object.assign({},u),{labels:t})})),s=(0,c.Fl)((()=>r.value.reduce((function(e,a){return e+a}),0))),r=(0,c.Fl)((()=>Object.values(n.value))),d={chart:{type:"line",animations:{enabled:!1},zoom:{enabled:!1}},dataLabels:{enabled:!1},stroke:{curve:"smooth"},title:{align:"center",margin:25},grid:{row:{colors:["#f3f3f3","transparent"],opacity:.5}},legend:{position:"bottom"},xaxis:{type:"datetime"},responsive:[{breakpoint:480,align:"center",options:{chart:{width:320,height:235},legend:{position:"bottom"}}}]},u={labels:o,title:{align:"center",margin:25},chart:{type:"pie",animations:{enabled:!0}},legend:{position:"bottom"},responsive:[{breakpoint:480,align:"center",options:{chart:{width:320},legend:{position:"bottom"}}}]};function v(e){var l;(null===(l=a.value)||void 0===l?void 0:l.players)&&(t=a.value.players.find((a=>a.id===e)),t&&(a.value.players=a.value.players.filter((a=>a.id!==e))))}function g(t){p.api.get(`/games/${t}`).then((e=>{a.value=e.data})).catch((a=>{var t,l;let i=a.message;(null===(t=a.response)||void 0===t?void 0:t.data)&&(i=null===(l=a.response)||void 0===l?void 0:l.data),e.notify({color:"negative",position:"top",message:i,icon:"report_problem"})}))}function f(t,l){var i;(null===(i=a.value)||void 0===i?void 0:i.id)&&p.api.post(`/games/${a.value.id}/players/${t}/stats/`,{date:l.toJSON(l)}).then((()=>{var e;(null===(e=a.value)||void 0===e?void 0:e.id)&&g(a.value.id)})).catch((a=>{var t,l;let i=a.message;(null===(t=a.response)||void 0===t?void 0:t.data)&&(i=null===(l=a.response)||void 0===l?void 0:l.data),e.notify({color:"negative",position:"top",message:i,icon:"report_problem"})}))}function h(t){var l;if(null===(l=a.value)||void 0===l?void 0:l.id){const l={name:t};p.api.post(`/games/${a.value.id}/players/`,l).then((()=>{var e;(null===(e=a.value)||void 0===e?void 0:e.id)&&g(a.value.id)})).catch((a=>{var t,l;let i=a.message;(null===(t=a.response)||void 0===t?void 0:t.data)&&(i=null===(l=a.response)||void 0===l?void 0:l.data),e.notify({color:"negative",position:"top",message:i,icon:"report_problem"})}))}}function w(l){var i;(null===(i=a.value)||void 0===i?void 0:i.id)&&p.api["delete"](`/games/${a.value.id}/players/${l}`).catch((l=>{var i,n,o;let s=l.message;(null===(i=l.response)||void 0===i?void 0:i.data)&&(s=null===(n=l.response)||void 0===n?void 0:n.data),e.notify({color:"negative",position:"top",message:s,icon:"report_problem"}),t&&(null===(o=a.value)||void 0===o||o.players.push(t))}))}function y(a,t){e.dialog({component:$,componentProps:{player:t,id:a}}).onOk((e=>{f(a,e)}))}function b(a,t){e.dialog({title:"Delete Player",message:`Are you sure to delete '${t}' player?`,cancel:!0,persistent:!0}).onOk((()=>{v(a),w(a)}))}function _(){e.dialog({title:"Add a new Player",message:"Player Name",prompt:{model:"",isValid:e=>e.length>1,type:"text"},cancel:!0}).onOk((e=>{h(e)}))}return{apiFetchGame:g,openAddWinnerDialog:y,deleteDialog:b,addPlayerDialog:_,gamePlayers:o,gamePlayersWins:n,gameSeries:r,pieChartOptions:u,lineChartOptions:d,nbGames:s,gamePlayersCumStats:i,gameDatesTs:l,game:a,chart:(0,c.iH)("pie")}},created(){this.apiFetchGame(parseInt(this.$route.params.id))}});var z=t(4379),A=t(2496),E=t(3269),G=t(3414),I=t(2035),T=t(1007);const K=(0,_.Z)(M,[["render",u],["__scopeId","data-v-0f4ad59e"]]),U=K;O()(M,"components",{QPage:z.Z,QCard:q.Z,QCardSection:k.Z,QTabs:A.Z,QTab:E.Z,QItem:G.Z,QItemSection:I.Z,QBtn:P.Z,QPageSticky:T.Z})}}]);