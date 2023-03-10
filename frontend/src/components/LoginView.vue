<template>
  <section>
    <b-checkbox v-model="hasError">Show errors</b-checkbox>
    <b-field
      label="Username"
      v-model="detalles.clues"
      :type="{ 'is-primary': hasError }"
      :message="{ 'Username is not available': hasError }"
    >
      <b-input maxlength="20"></b-input>
    </b-field>

    <b-field
      label="Password"
      v-model="detalles.password"
      :type="{ 'is-primary': hasError }"
      :message="[
        { 'Password is too short': hasError },
        { 'Password must have at least 8 characters': hasError },
      ]"
    >
      <b-input type="password" maxlength="20"></b-input>
    </b-field>
    <b-field horizontal
      ><!-- Label left empty for spacing -->
      <p class="control">
        <b-button label="Login" type="is-primary" @click="submit()" />
      </p>
    </b-field>
  </section>
</template>

<script>
import Vue from "vue";
import Router from "vue-router";
import DialogosService from "../services/DialogosService";
import PuerperaService from "../services/PuerperaService";
Vue.use(Router);
const router = new Router();
export default {
   
  data: () => ({
    hasError:false,
    detalles: {
      clues: "maria",
      password: "kfc",
    },
    router,
   
  }),
  methods: {
    async submit() {
      
      const url = "http://localhost:5000/auth/login"
      const res = await fetch(url,  {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        credentials: "include",
        body: JSON.stringify(this.detalles)
      })
      const status = res.status
      if (status == 200) {
        console.log("200 ok")
        this.router.push("/agregar-embarazada")
      } else {
        console.log("no auth")
      }
    },
    async login() {
      if (!this.detalles.clues == null) {
        return DialogosService.mostrarNotificacionError("Campos vacios");
      }
      const cargaUtil = {
        clues: this.detalles.clues,
        password: this.detalles.password
      };

      const respuesta = await PuerperaService.login(cargaUtil);
      if (respuesta != true) {
        console.log("respuesta save", respuesta);
         console.log("no auth",this.detalles);
        DialogosService.mostrarError("registro con errores");
      } else {

        DialogosService.mostrarExitoso(
          "Datos guardados correctament en la base de datos"
        );
   
      }
    },


  },
};
</script>