<template>
  <div id="form">
    <section>
      <b-field label="CLUES" v-model="elementos.clues">
        <b-input placeholder="CLUES"></b-input>
      </b-field>

      <b-field
        label="Nombre Unidad"
        type="is-danger"
        message="Escriba el nombre de la unidad"
        v-model="elementos.nombreUnidad"
      >
        <b-input type="text" maxlength="30"> </b-input>
      </b-field>

      <b-field
        label="Nombre de Usuario"
        type="is-success"
        message="Este usuario e..."
        v-model="elementos.user"
      >
        <b-input maxlength="30"></b-input>
      </b-field>

      <b-field
        label="Contarseña"
        type="is-danger"
        :message="[
          'Password is too short',
          'Password must have at least 8 characters',
        ]"
        v-model="elementos.password"
      >
        <b-input type="password" maxlength="30"></b-input>
      </b-field>
      <button> hola</button>
    </section>
  </div>
</template>
<script>
import UserService from "../services/UserService";
import DialogosService from "../services/DialogosService";
export default {
  data: () => ({
    elementos: {
      clues: "",
      nombreUnidad: "",
      user: "",
      password: "",
    },
  }),
  methods: {
    async saveUser() {
      const cargaUtil = {
        clues: this.elementos.clues,
      };
      const respuesta = await UserService.insertUser(cargaUtil);
      if (respuesta != true) {
        console.log("respuesta save", respuesta);
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
<style>
#form {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-bottom: 10px;
}
</style>