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
        <div class="navbar-item has-dropdown is-hoverable">
          <a class="navbar-link"> Usuario </a>
          <div class="navbar-dropdown">
            <b-navbar-item tag="router-link" :to="{ name: 'Login' }">
              Perfil</b-navbar-item>
            <b-navbar-item tag="router-link" :to="{ name: 'Login' }"> Configurar</b-navbar-item>
            <b-navbar-item tag="router-link" :to="{ name: 'Login' }">Cerrar sesion </b-navbar-item>

          </div>
        </div>

        <div class="navbar-item has-dropdown is-hoverable">
          <a class="navbar-link"> Embarazada </a>
          <div class="navbar-dropdown">
            <b-navbar-item tag="router-link" :to="{ name: 'RegistrarEmbarazada' }">
              Registrar Embarazada</b-navbar-item>
            <b-navbar-item tag="router-link" :to="{ name: 'RegistrarPuerpera' }"> Acturalizar datos</b-navbar-item>
            <b-navbar-item tag="router-link" :to="{ name: 'RegistrarPuerpera' }"> Mostrar lista</b-navbar-item>
          </div>
        </div>

        <div class="navbar-item has-dropdown is-hoverable">
          <a class="navbar-link"> Puerpera </a>
          <div class="navbar-dropdown">
            <b-navbar-item tag="router-link" :to="{ name: 'RegistrarPuerpera' }">
              Registrar Puerpera</b-navbar-item>
            <b-navbar-item tag="router-link" :to="{ name: 'RegistrarPuerpera' }"> Acturalizar datos</b-navbar-item>
            <b-navbar-item tag="router-link" :to="{ name: 'RegistrarPuerpera' }"> Mostrar lista</b-navbar-item>
          </div>
        </div>
        <div class="navbar-item has-dropdown is-hoverable">
          <a class="navbar-link"> Informes </a>
          <div class="navbar-dropdown">
            <b-navbar-item tag="router-link" :to="{ name: 'RegistrarPuerpera' }">
              Region</b-navbar-item>
            <b-navbar-item tag="router-link" :to="{ name: 'RegistrarPuerpera' }"> Municipio</b-navbar-item>
            <b-navbar-item tag="router-link" :to="{ name: 'RegistrarPuerpera' }"> Localidad</b-navbar-item>
          </div>
        </div>
        <div class="navbar-item has-dropdown is-hoverable">
          <a class="navbar-link"> Tarjetas </a>
          <div class="navbar-dropdown">
            <b-navbar-item tag="router-link" :to="{ name: 'RegistrarPuerpera' }">
              Region</b-navbar-item>
            <b-navbar-item tag="router-link" :to="{ name: 'RegistrarPuerpera' }"> Municipio</b-navbar-item>
            <b-navbar-item tag="router-link" :to="{ name: 'RegistrarPuerpera' }"> Localidad</b-navbar-item>
          </div>
        </div>


      </template>

      <template #end>
        <div class="navbar-item has-dropdown is-hoverable" >
          <a class="navbar-link"  > Soporte  </a>
          
          <div class="navbar-dropdown" >
            <b-navbar-item tag="router-link" :to="{ name: 'RegistrarPuerpera' }">
              Acerca de </b-navbar-item>
            <b-navbar-item tag="router-link" :to="{ name: 'RegistrarPuerpera' }"> Ayuda</b-navbar-item>
          </div>
        </div>
      </template>
    </b-navbar>
    <section class="section">
      <router-view></router-view>
    </section>
    <footer class="footer">
      <div class="content has-text-centered">
       
        <p>
      <a> Sistema integral de salud materna</a> © 2022    </p>
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