
<template>
  <div>
    <b-navbar type="is-info">
      <template #brand>
        <b-navbar-item tag="router-link" :to="{ path: '/' }"> </b-navbar-item>
      </template>
      <template #start>
        <b-navbar-item tag="router-link" :to="{ name: 'Mujer' }"
          ><b-icon icon="home" class="mr-1"></b-icon>Inicio</b-navbar-item
        >
        <div class="navbar-item has-dropdown is-hoverable">
          <a class="navbar-link"> Usuario </a>
          <div class="navbar-dropdown">
            <b-navbar-item tag="router-link" :to="{ name: 'Login' }">
              Perfil</b-navbar-item
            >
            <b-navbar-item tag="router-link" :to="{ name: 'Login' }">
              Configurar</b-navbar-item
            >
            <b-navbar-item tag="router-link" :to="{ name: 'Login' }"
              >Cerrar sesion
            </b-navbar-item>
          </div>
        </div>

        <div class="navbar-item has-dropdown is-hoverable">
          <a class="navbar-link"> Embarazada </a>
          <div class="navbar-dropdown">
            <b-navbar-item
              tag="router-link"
              :to="{ name: 'RegistrarEmbarazada' }"
            >
              Registrar Embarazada</b-navbar-item
            >
            <b-navbar-item
              tag="router-link"
              :to="{ name: 'RegistrarPuerpera' }"
            >
              Acturalizar datos</b-navbar-item
            >
            <b-navbar-item
              tag="router-link"
              :to="{ name: 'RegistrarPuerpera' }"
            >
              Mostrar lista</b-navbar-item
            >
          </div>
        </div>

        <div class="navbar-item has-dropdown is-hoverable">
          <a class="navbar-link"> Puerpera </a>
          <div class="navbar-dropdown">
            <b-navbar-item
              tag="router-link"
              :to="{ name: 'RegistrarPuerpera' }"
            >
              Registrar Puerpera</b-navbar-item
            >
            <b-navbar-item
              tag="router-link"
              :to="{ name: 'RegistrarPuerpera' }"
            >
              Acturalizar datos</b-navbar-item
            >
            <b-navbar-item
              tag="router-link"
              :to="{ name: 'RegistrarPuerpera' }"
            >
              Mostrar lista</b-navbar-item
            >
          </div>
        </div>
        <div class="navbar-item has-dropdown is-hoverable">
          <a class="navbar-link"> Informes </a>
          <div class="navbar-dropdown">
            <b-navbar-item tag="router-link" :to="{ name: 'Informes' }">
              Region</b-navbar-item
            >
            <b-navbar-item
              tag="router-link"
              :to="{ name: 'RegistrarPuerpera' }"
            >
              Municipio</b-navbar-item
            >
            <b-navbar-item
              tag="router-link"
              :to="{ name: 'RegistrarPuerpera' }"
            >
              Localidad</b-navbar-item
            >
          </div>
        </div>
        <div class="navbar-item has-dropdown is-hoverable">
          <a class="navbar-link"> Tarjetas </a>
          <div class="navbar-dropdown">
            <b-navbar-item
              tag="router-link"
              :to="{ name: 'RegistrarPuerpera' }"
            >
              Region</b-navbar-item
            >
            <b-navbar-item
              tag="router-link"
              :to="{ name: 'RegistrarPuerpera' }"
            >
              Municipio</b-navbar-item
            >
            <b-navbar-item
              tag="router-link"
              :to="{ name: 'RegistrarPuerpera' }"
            >
              Localidad</b-navbar-item
            >
          </div>
        </div>
      </template>

      <template #end>
        <div class="navbar-item has-dropdown is-hoverable">
          <a class="navbar-link"> Soporte </a>

          <div class="navbar-dropdown">
            <b-navbar-item
              tag="router-link"
              :to="{ name: 'RegistrarPuerpera' }"
            >
              Acerca de
            </b-navbar-item>
            <b-navbar-item
              tag="router-link"
              :to="{ name: 'RegistrarPuerpera' }"
            >
              Ayuda</b-navbar-item
            >
          </div>
        </div>
      </template>
    </b-navbar>
    <section class="section">
      <router-view> </router-view>
    </section>
    <footer class="footer">
      <div class="content has-text-centered">
        <div class="wrapper">
          <div>
            <img :src="secre" />
          </div>
          <div>
            <img :src="materno" />
          </div>
          <div>
            <img :src="tec" />
          </div>
        </div>
        <p><a> Sistema integral de salud materna</a> © 2022</p>
      </div>
    </footer>
  </div>
</template>

<script>
import Logo from "./assets/parzibyte.png";
import materna from "./assets/materno.png";
import secretaria from "./assets/secretaria.png";
import imgtec from "./assets/tec.png";
export default {
  name: "app",
  data: () => ({
    logo: Logo,
    materno: materna,
    secre: secretaria,
    tec: imgtec,
  }),
  methods: {},
};
</script>

<style>
#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  background-color: #182a3f;
  min-height: 100vh;
}
.wrapper {
  display: grid;
  grid-template-columns: 1fr .5fr .5fr;
}
img {
  display: block;
  margin: auto;
  min-height: 5vh;
  width: 45%;
  height: 30;
}

#building {
  background-color: #fafffe;
  background-size: initial;
  background-position: left left;
  background-repeat: no-repeat;
  min-height: 5vh;
  width: 10%;
  height: 8%;
}
</style>