<template>
  <div class="app">
    <nav class="nav">
      <ul class="menu">
        <li class="item">
      <router-link class="nav-link" to="/">Processi terminati</router-link></li>
      <li class="item">
      <router-link class="nav-link" to="/process">Gestione processo</router-link></li>
      <li class="item">
      <router-link class="nav-link" to="/test-page">Test funzioni</router-link></li>
      <li class="item">
      <router-link class="nav-link" to="/configuration">Configurazione</router-link></li>
      </ul>
    </nav>
    <router-view class="content"/>
  </div>
</template>

<style>
.app{
  margin: 0;
  font: small-caps bold 14px/1 sans-serif;
}
.nav{
background: #222;
  padding: 0 15px;
  
}
.nav-link{
  color: white;
  text-decoration: none;
  
}
.menu{   
    list-style-type: none;
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
    align-items: center;
    order: 1;
    align-items: flex-start;     
    flex-wrap: nowrap;
    background: none;
    
}
.menu li a {
    display: block;
    padding: 15px 5px;
}
.item {
    padding: 10px;
    position: relative;
    display: block; 
    width: auto;
    order: 1;
}
.item.button {
    padding: 9px 5px;
}
.item:not(.button) a:hover,
.item a:hover::after {
    color: #ccc;
}
.content{
  margin: 0 10% 0 10% ; 
  
}
.content button{
  border-radius: 10%;
  padding: 10pt;
  background-color: gray;
  color: white;
  margin: 5pt;
}
table {
  border-collapse: collapse;
  width: 100%;
}

td, th {
  border: 1px solid #ddd;
  padding: 8px;
}

tr:nth-child(even){background-color: #f2f2f2;}



th {
  padding-top: 12px;
  padding-bottom: 12px;
  text-align: left;
  background-color: grey;
  color: white;
}

</style>
