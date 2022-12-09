<template>
  <section>


    <b-table  class="table is-bordered is-striped is-narrow is-hoverable is-fullwidth" :data="data" >
      <template v-for="column in columns">
        <b-table-column .key="column.id" v-bind="column">
          <template 
          v-if="column.searchable && !column.numeric" 
          #searchable="props" >
            <b-input 
            v-model="props.filters[props.column.field.nombre]"
              placeholder="Buscar" 
              icon="magnify"
               size="is-small" />
          </template>
          <template v-slot="props">
            {{ props.row[column.field.id] }}
          </template>
        </b-table-column>
      </template>
    </b-table>
  </section>

</template>

<script>
import VehiculosService from "../services/VehiculosService";

export default {
  data() {
    return {
      data: [],
      cargando: false,
      columns: [
        {
          field: 'id',
          label: 'No. Progreso',
          width: '100',
          numeric: true,
          searchable: true,
          centered: true
        },
        {
          field: 'noExpediente',
          label: 'Nombre completo',
          searchable: true,
          centered: true
        },
        {
          field: 'nombre',
          label: 'SDG',
          centered: true
        },
        {
          field: 'curp',
          label: 'Fecha de consutal',
          searchable: true,
          centered: true
        },
        {
          field: 'telefono',
          label: 'telefono',
          searchable: true,
          centered: true
        }
      
      ],
      async mounted() {
  
    await this.obtenerVehiculos();
  },
      methods:{
        async obtenerEmbarazadas() {
      this.cargando = true;
          this.embarazada = await VehiculosService.obtenerVehiculos();
          console.log(this.embarazada);
      this.cargando = false;
      }
      
    },
      
    }
  }
}
</script>
<style>
#tabla {
  justify-content: center;
}
</style> 