<template>


  <article class="panel is-info">
    <p class="panel-heading">
      Bucar embarazada
    </p>



    <section>

      <b-field>
        <div class="control">
          <b-button label="Eliminar paciente seleccionado" type="is-danger" icon-left="delete-forever"
            :disabled="!selected" @click="selected = null" />
        </div>


        <b-button label="Modificar paciente seleccionado" type="is-success" icon-left="border-color"
          :disabled="!selected" @click="selected = null" />



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
        <div class="control">
          <b-switch v-model="dateSearchable">Fecha de nacimiento</b-switch>
        </div>


      </b-field>
      <b-table :data="vehiculos"
          :loading="cargando"
          :mobile-cards="true"
          hoverable>
        <b-table-column
            searchable
            field="descripcion"
            label="Descripción"
            v-slot="props"
            sortable
          >
            {{ props.row.nombre}}
          
          </b-table-column>
      </b-table>
      
      <br />

    </section>
    <template>
    <b-table :data="date" :columns="columnss">
      
    </b-table>
</template>
  </article>

</template>


<script>
import VehiculosService from "../services/VehiculosService";
export default {
  
  data() {
   
    return {
      selected: false,
      vehiculos:[],
      cargando: false,
      date: [
        {
          id: 1,
          user: { first_name: "HortenciaGarciaMendez", last_name: "HASR191098MCHLLC07" },
          date: "10/15/2016",
          gender: "27 años"
        },
        {
          id: 2,
          user: { first_name: "John", last_name: "Jacobs" },
          date: "12/12/2022 06:00:53",
          gender: "Male"
        },
        {
          id: 3,
          user: { first_name: "Tina", last_name: "Gilbert" },
          date: "2016/04/26 06:26:28",
          gender: "Female"
        },
        {
          id: 4,
          user: { first_name: "Clarence", last_name: "Flores" },
          date: "2016/04/10 10:28:46",
          gender: "Male"
        },
        {
          id: 5,
          user: { first_name: "Anne", last_name: "Lee" },
          date: "2016/12/06 14:38:38",
          gender: "Female"
        },
        {
          id: 6,
          user: { first_name: "Sara", last_name: "Armstrong" },
          date: "2016/09/23 18:50:04",
          gender: "Female"
        },
        {
          id: 7,
          user: { first_name: "Anthony", last_name: "Webb" },
          date: "2016/08/30 23:49:38",
          gender: "Male"
        },
        {
          id: 8,
          user: { first_name: "Andrew", last_name: "Greene" },
          date: "2016/11/20 14:57:47",
          gender: "Male"
        },
        {
          id: 9,
          user: { first_name: "Russell", last_name: "White" },
          date: "2016/07/13 09:29:49",
          gender: "Male"
        },
        {
          id: 10,
          user: { first_name: "Lori", last_name: "Hunter" },
          date: "2016/12/09 01:44:05",
          gender: "Female"
        }
      ],
      stickyHeaders: true,
      dateCurp: false,
      dateNombre: false,
      dateSearchable: false

    };

  },

  computed: {
    columnss() {
      return [

        {
          field: "id",
          label: "No. Progreso",
          width: "40",
          numeric: true,
          sticky: true,
          headerClass: "is-sticky-column-one",
          cellClass: "is-sticky-column-one",
          centered: true,


        },
        {
          field: "user.last_name",
          label: "Curp",
          searchable: this.dateCurp,
          centered: true,
          sticky: true,

        },
        {
          field: "user.first_name",
          label: "Nombre",
          searchable: this.dateNombre,
          centered: true,
          sticky: true,


        },

        {
          field: "date",
          label: "Fecha de nacimiento",
          searchable: this.dateSearchable,
          centered: true,
          sticky: true

        },
        {
          field: "gender",
          label: "Edad",
          centered: true,
          width: "50"
        },
        {
          field: "id2",
          label: "Consulta de 7 días",
          centered: true,
          width: "120",
          height: "60"
        },
        {
          field: "id3",
          label: "Consulta de 28 días",
          centered: true,
          width: "50"
        },
        {
          field: "id4",
          label: "Consulta de 40 días",
          centered: true,
          width: "50"
        },
        {
          field: "id5",
          label: "Signos de alarma",
          centered: true,
          width: "50"
        },
        {
          field: "id6",
          label: "Atención de parto",
          centered: true,
          width: "50"

        },
        {
          field: "id7",
          label: "Lugar del parto",
          centered: true,
          width: "50"
        },
        {
          field: "id8",
          label: "Resolución del embarazo",
          centered: true,
          width: "50"
        },
        {
          field: "id9",
          label: "No. de producto",
          centered: true,
          width: "50"
        },
        {
          field: "id11",
          label: "APEO",
          centered: true,
          width: "50"
        },
        {
          field: "id12",
          label: "Puerpera aceptante",
          centered: true,
          width: "50"
        },
        {
          field: "id13",
          label: "Consulta pregestacional",
          centered: true,
          width: "50"
        },
        {
          field: "id14",
          label: "Fecha de parto",
          centered: true,
          width: "50"
        },
        {
          field: "id15",
          label: "Derechohabiencia",
          centered: true,
          width: "50",
          height: "60"
        }
      ];
    },
    checkableColumns() {
      return [
        {
          field: "id",
          label: "ID",
          width: "40",
          numeric: true,
          sticky: false,
        },
        {
          field: "user.first_name",
          label: "Nombre",
          searchable: this.dateNombre,
          centered: true,
          sticky: false
        },
        {
          field: "user.last_name",
          label: "Curp",
          searchable: this.dateCurp,
          centered: true,
          sticky: false
        },
        {
          field: "date",
          label: "Date",
          searchable: this.dateSearchable,
          centered: true,
          sticky: false,
        },
        {
          field: "gender",
          label: "Gender",
          width: "50"
        },
        {
          field: "id",
          label: "Column A",

        },
        {
          field: "id",
          label: "Column B"
        },
        {
          field: "id",
          label: "Column C"
        },
        {
          field: "id",
          label: "Column D"
        },
        {
          field: "id",
          label: "Column E"
        },
        {
          field: "id",
          label: "Column F"
        },
        {
          field: "id",
          label: "Column G"
        },
        {
          field: "id",
          label: "Column H"
        },
        {
          field: "id",
          label: "Column I"
        },
        {
          field: "id",
          label: "Column L"
        },
        {
          field: "id",
          label: "Column M"
        },
        {
          field: "id",
          label: "Column N"
        },
        {
          field: "id",
          label: "Column O"
        }
      ];
    },

  },
  async mounted() {
 
    await this.obtenerVehiculos();
  },
  methods:
  {
    async obtenerVehiculos() {
      this.cargando = true;
      
      this.vehiculos = await VehiculosService.obtenerVehiculos();
       console.log(this.vehiculos);
      this.cargando = false;
    },
  }
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