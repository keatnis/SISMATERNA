<template>
  <div>
    <b-navbar type="is-info">
      <template #brand>
        <b-navbar-item tag="router-link" :to="{ path: '/' }">
          <img :src="logo" />
        </b-navbar-item>
      </template>
      <template #start>
        <b-navbar-item tag="router-link" :to="{ name: 'Mujer' }"
          ><b-icon icon="home" class="mr-1"></b-icon>Inicio</b-navbar-item
        >
        <b-navbar-item tag="router-link" :to="{ name: 'Informes' }"
          ><b-icon icon="table-check" class="mr-1"></b-icon>
         Informes</b-navbar-item
        >
        
        <b-navbar-item tag="router-link" :to="{ name: 'Ajustes' }"
          ><b-icon icon="cog-outline" class="mr-1"></b-icon>
          Ajustes</b-navbar-item
        >
        <div class="navbar-item has-dropdown is-hoverable">
          <a class="navbar-link"> Usuario </a>
          <div class="navbar-dropdown">
          <b-navbar-item tag="router-link" :to="{ name: 'Login' }"> 
            Perfil</b-navbar-item>
            <b-navbar-item tag="router-link" :to="{ name: 'Login' }"> Configurar</b-navbar-item>
            <b-navbar-item tag="router-link" :to="{ name: 'Login' }">Cerrar sesion </b-navbar-item>
          
          </div>
        </div>
        
      </template>

      <template #end>
        <b-navbar-item tag="div">
          <div class="buttons">
            <a href="https://parzibyte.me/#contacto" class="button is-danger">
              <b-icon icon="help"></b-icon>
              <strong>Ayuda y soporte</strong>
            </a>
          </div>
        </b-navbar-item>
      </template>
    </b-navbar>
    <section class="section">
      <router-view></router-view>
    </section>
    <footer class="footer">
      <div class="content has-text-centered">
        <p>
          <strong>Sistema Integral de Salud Materna</strong>
           por <a href="https://itsm-tlapa.me/blog">ITSM</a>
        </p>
      </div>
    </footer>
  </div>
</template>

<script>
import Logo from "./assets/parzibyte.png";
export default {
  name: "app",
  data: () => ({
    logo: Logo,
  }),
  methods: {},
};
</script>