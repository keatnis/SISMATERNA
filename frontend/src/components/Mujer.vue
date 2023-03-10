<template>
  <section>
    <b-table
      :data="puerperas"
      :columns="columns"
      :sticky-header="stickyHeaders"
    ></b-table>
  </section>
</template>

<script>
import DialogosService from "../services/DialogosService";

import PuerperaService from "../services/PuerperaService";
export default {
  data() {
    return {
      stickyHeaders: true,
      puerperas: [],
      columns: [
        {
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
          label: "Pregestacional",
        },
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
        DialogosService.mostrarNotificacionError(
          "No se puede obtener la lista de puerperas"
        );
      }
    },
  },
};
</script>
<style>
#tabla {
  justify-content: center;
}
</style>
