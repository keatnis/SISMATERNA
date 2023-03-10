<template>
  <section>
    <article class="panel is-danger centered ">
    <p class="panel-heading " >
      Proximas citas:
    </p></article>
    
    <b-table :data="embarazadas" :columns="columns" :sticky-header="stickyHeaders" width=700></b-table>  
   


  </section>
</template>

<script>

import DialogosService from "../services/DialogosService";

import EmbarazadaService from "../services/EmbarazadaService";
export default {
  data() {
    return {

      stickyHeaders: true,
      dateCurp: false,
      dateNombre: false,
      dateSearchable: false,
      embarazadas: [],
      columns: [
        {
          field: "id",
          label: "NP",
          width: "25",
          
          numeric: true,
          
        },
        {
          field: "curp",
          label: "Curp",
          centered: true,
          searchable: this.datecurp,

        },
       

        {
          field: "noExpediente",
          label: "No. de expediente clinico",
          width: "70",
          numeric: true,
          sticky: true,
          
          centered: true,

         
        },
        {
          field: "nombre",
          label: "Nombre de la embarazada",
          numeric: true,
          sticky: true,
          
          centered: true,

          
        },

        
        {
          field: "telefono",
          label: "Numero telefonico",
          centered: true,

        },
        {
          field: "SGA",
          label: "SDG Actual",
          centered: true,
          
        },
        {
          field: "FechaConsulta",
          label: "Fecha de consulta",
          centered: true,
          
        },
        
      ],
    };
  },

  async mounted() {
    await this.GetEmbarazada();
  },
  methods: {

    async GetEmbarazada() {

      try {
        this.embarazadas = await EmbarazadaService.GetEmbarazada();
      } catch {
        DialogosService.mostrarNotificacionError("No se puede obtener la lista de embarazada");
      }
    }
  },
};


</script>
<style>
#tabla {
  justify-content: center;
}
</style>
