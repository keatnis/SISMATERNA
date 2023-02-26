<template>
  <section>
    <article class="panel is-info">
      <p class="panel-heading">
        Bucar puerpera:
      </p>
    </article>

    <b-field>
      <div class="control">
        <b-button label="Eliminar paciente seleccionado" type="is-danger" icon-left="delete-forever" :disabled="!selected"
          @click="selected = null" />
      </div>


      <b-button label="Modificar paciente seleccionado" type="is-success" icon-left="border-color" :disabled="!selected"
        @click="selected = null" />
    </b-field>


    <b-field grouped group-multiline>
      <div class="control">
        <b-switch v-model="stickyHeaders">Encabezado fijo</b-switch>
      </div>
      <div class="control">
        <b-switch v-model="dateCurp">CURP</b-switch>
      </div>
      <div class="control">
        <b-switch v-model="dateNombre"> Nombre completo</b-switch>
      </div>



    </b-field>


    <b-table :data="puerperas" :columns="columns" :sticky-header="stickyHeaders"></b-table>





  </section>
</template>


<script>

import DialogosService from "../services/DialogosService";

import PuerperaService from "../services/PuerperaService";
export default {
  data() {
    return {

      stickyHeaders: true,
      dateCurp: false,
      dateNombre: false,
      dateSearchable: false,
      embarazadas: [],
      localidades: [],
      selectedOptions: [],
      puerperas: [],
      columns: [{
        field: "NP",
        label: "ID",
        width: "40",
        numeric: true,

      },
      {
        field: "consultasiete",
        label: "Consulta 7",
      },
      {
        field: "consultaveinte",
        label: "Consulta 20",
      },

      {
        field: "consultacuarenta",
        label: "Consulta 40",
        centered: true,
      },

      {
        field: "signos",
        label: "Signos",
      },
      {
        field: "atencionparto",
        label: "Atendio Parto",

      },
      {
        field: "lugarparto",
        label: "Lugar Parto",

      },
      {
        field: "resolucion",
        label: "Resolucion",

      },
      {
        field: "producto",
        label: "Producto",

      },
      {
        field: "aceptante",
        label: "Aceptance",

      },
      {
        field: "pregestacional",
        label: "Pregestacional"
      }
      ],
    };
  },
  async mounted() {
    await this.ObtenerPuerperas();
  },
  methods: {

    async ObtenerPuerperas() {

      try {
        this.puerperas = await PuerperaService.obtenerPuerpera();
      } catch {
        DialogosService.mostrarNotificacionError("No se puede obtener la lista de puerperas");
      }
    }
  },
};

</script>

<style>
.is-sticky-column-one {
  background: #7cc4eb !important;
  color: white !important;
}

.is-sticky-column-two {
  background: #6ca8ee !important;
  color: white !important;
}
</style>