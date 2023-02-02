<template>
  <section>

        <b-table
          :data="vehiculos"
        
          
          hoverable
        >
          <b-table-column
            searchable
            field="descripcion"
            label="Descripción"
            v-slot="props"
            sortable
          >
            {{ props.row.id}}
          </b-table-column>
          <b-table-column
            searchable
            field="placas"
            label="Placas"
            sortable
            v-slot="props"
          >
            {{ props.row.nombre }}
          </b-table-column>
          <b-table-column
            searchable
            field="propietario"
            label="Propietario"
            sortable
            v-slot="props"
          >
            {{ props.row.curp}}
          </b-table-column>


          <b-table-column
            field="fechaEntrada"
            label="Entrada"
            v-slot="props"
            sortable
          >
            {{ props.row.curp }}
          </b-table-column>
      
          <template #empty>
            <div class="has-text-centered">No hay registros</div>
          </template>
        </b-table>
   <div>
    {{vehiculos}}
   </div>
    
  </section>
</template>

<script>

export default {
  data() {
    return {
      vehiculos: [],
      columns: [
      {
          field: "id",
          label: "No. Progreso",
          width: "100",
          numeric: true,
          searchable: true,
          centered: true,
        },
        {
          field: "noExpediente",
          label: "Nombre completo",
          searchable: true,
          centered: true,
        },
        {
          field: "nombre",
          label: "SDG",
          centered: true,
        },
        {
          field: "curp",
          label: "Fecha de consutal",
          searchable: true,
          centered: true,
        },
        {
          field: "telefono",
          label: "telefono",
          searchable: true,
          centered: true,
        },
    ]
  }
},

};
</script>
<style>
#tabla {
  justify-content: center;
}
</style> 