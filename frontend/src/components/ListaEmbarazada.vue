<template>
  <section>
    <hero-bar> Listado de Embarazadas </hero-bar>
    <hr />
    <card-component title="Listado de Embarazadas">
      <div class="notification is-info is-light">
        <b-button
          type="is-success"
          slot="right"
          @click="exportExcel()"
          class="button"
        >
          Exportar a Excel
        </b-button>
      </div>

      <b-field grouped group-multiline>
        <div class="control">
          <b-switch v-model="dateCurp">CURP</b-switch>
        </div>
        <div class="control">
          <b-switch v-model="dateNombre"> Nombre completo</b-switch>
        </div>
      </b-field>
      <b-table
        class="table"
        :data="false ? [] : embarazadas"
        :columns="columns"
        :bordered="true"
        :striped="true"
        :narrowed="true"
        :hoverable="true"
        :scrollable="true"
        :loading="embarazadas.length > 0 ? false : true"
      >
      </b-table>
    </card-component>
  </section>
</template>

<script>
import DialogosService from "../services/DialogosService";
import HeroBar from "@/components/HeroBar.vue";
import EmbarazadaService from "../services/EmbarazadaService";
import CardComponent from "./CardComponent.vue";
// import readXlsFile from "read-excelfile"
//import exportXlsFile from "export-from-json"

export default {
  name: "ListaEmbarazada",
  components: {
    HeroBar,
    CardComponent,
  },

  data() {
    return {
      checkedRows: [],
      isModalActive: false,
      trashObject: null,

      stickyHeaders: true,
      isbordered: true,
      dateCurp: false,
      dateNombre: false,
      embarazadas: [],

      columns: [
        {
          field: "id",
          label: "NP",
          width: "25",
          numeric: true,

          sortable: true,
        },
        {
          field: "curp",
          label: "Curp",

          searchable: this.datecurp,
        },
        {
          field: "nuevoingreso",
          label: "Nuevo ingreso",
        },

        {
          field: "noExpediente",
          label: "No. Exp",
        },
        {
          field: "nombre",
          label: "Nombre_completo",
          numeric: true,
          sticky: true,

          searchable: this.dateNombre,
          headerClass: "nombretb",
          cellClass: "celdanm",
        },

        {
          field: "FechaNacimiento",
          label: "Fecha_de_nacimiento ",
        },

        {
          field: "edad",
          label: "Edad",
        },
        {
          field: "domicilioReferencia",
          label: "Direccion o referencia ",
        },

        {
          field: "localidad",
          label: "Localidad",
          width: "180",
        },
        {
          field: "municipio",
          label: "Municipio",
        },
        {
          field: "telefono",
          label: "Numero telefonico",
        },
        {
          field: "lenguaIndigena",
          label: "Habla lengua indigena",
        },
        {
          field: "emigro",
          label: "Emigro",
        },
        {
          field: "derechohabiencia",
          label: "Derechohabiencia",
        },
        {
          field: "violencia",
          label: "Violencia",
        },
        {
          field: "detecciones",
          label: "Detecciones",
        },

        {
          field: "ConsultaPregestacional",
          label: "Consulta de riesgo pregestacional",
        },

        {
          field: "comorbilidades",
          label: "Comorbilidades",
        },

        {
          field: "gestas",
          label: "Gestas",
        },
        {
          field: "paras",
          label: "Paras",
        },
        {
          field: "abortos",
          label: "Abortos",
        },
        {
          field: "cesareas",
          label: "Cesareas",
        },

        {
          field: "FechaUltimoEvento",
          label: "Fecha ultimo parto, cesarea o aborto",
        },
        {
          field: "presentComplicaciones",
          label: "¿Presentó complicaciones en el embarazo previo?",
        },
        {
          field: "complicaciones",
          label: "¿Que complicaciones presento?",
        },
        {
          field: "FechaUlmaMenstruacion",
          label: "Fecha ultima menstruación",
        },
        {
          field: "FechaProbableParto",
          label: "Fecha probable de parto",
        },
        {
          field: "trimestre",
          label: "Trimestre actual",
        },
        {
          field: "SGA",
          label: "SDG Actual",
        },
        {
          field: "FechaConsulta",
          label: "Fecha de consulta",
        },
        {
          field: "noConsultasEmbarazo",
          label: "No.de consulta durante el embarazo",
        },
        {
          field: "noConsultasMes",
          label: "No. consulta en el mes",
        },
        {
          field: "abortos",
          label: "Abortos",
        },
        {
          field: "rubeola",
          label: "Rubeola",
        },
        {
          field: "FechaVacunaTDPrimera",
          label: "Fecha TD primera",
        },
        {
          field: "FechaVacunaTDSegunda",
          label: "Fecha TD segunda",
        },
        {
          field: "FechaVacunaTDRefuerzo",
          label: "Fecha TD refuerzo",
        },
        {
          field: "FechaVacunaTDPA",
          label: "Fecha TDPA",
        },
        {
          field: "FechaVacunaInfluenza",
          label: "Fecha influenza estacional ",
        },
        {
          field: "vacunaCOVID",
          label: "AntiCovid-19",
        },
        {
          field: "grupoRH",
          label: "Grupo y RH",
        },
        {
          field: "ego",
          label: "EGO",
        },
        {
          field: "BiometriaHematica",
          label: "Biometria ematica",
        },
        {
          field: "leucocitos",
          label: "Leucocitos",
        },
        {
          field: "plaquetas",
          label: "Plaquetas",
        },
        {
          field: "vdrl",
          label: "VDRL",
        },
        {
          field: "vih",
          label: "VIH",
        },
        {
          field: "estadoGlucosa",
          label: "estadoGlucosa",
        },
        {
          field: "resultadoGlucosa",
          label: "resultadoGlucosa",
        },
        {
          field: "caracteristicasFetls",
          label: "Caracteristicas fetales",
        },
        {
          field: "malformaciones",
          label: "Malformaciones",
        },
        {
          field: "LiquidoAmiotico",
          label: "Liquido amiotico",
        },
        {
          field: "placenta",
          label: "Placenta",
        },
        {
          field: "FechaProbableUSG",
          label: "Fecha Probable de parto por USG",
        },
        {
          field: "referencia",
          label: "Referencia",
        },
        {
          field: "acudio_refe",
          label: "Acudio a refenrecia",
        },
        {
          field: "contrareferencia",
          label: "Contrareferencia recibida",
        },
        {
          field: "planSeguridad",
          label: "¿Se_le_realizo_plan_de seguridad?",
        },
        {
          field: "lugarParto",
          label: "¿Donde atendera parto?",
        },
        {
          field: "quienAtenderaParto",
          label: "¿Quien atendera parto?",
        },
        {
          field: "transporteAME",
          label: "Cuenta con transporte AME",
        },
        {
          field: "FechaEvento",
          label: "Fecha de vento",
        },
        {
          field: "observaciones",
          label: "Observaciones",
        },
        {
          field: "diagnostico",
          label: "Diagnostico",
        },
      ],
    };
  },
  computed: {
    paginated() {
      return this.embarazadas.length > this.perPage;
    },
  },
  async mounted() {
    await this.GetEmbarazada();
  },

  methods: {
    exportExcel: function () {
      var XLSX = require("xlsx");
      var hoy = new Date();
      let data = XLSX.utils.json_to_sheet(this.embarazadas);
      const workbook = XLSX.utils.book_new();

      const filename = "censo_" + hoy.getDay() + "/" + hoy.getMonth();
      XLSX.utils.book_append_sheet(workbook, data, filename);
      XLSX.writeFile(workbook, `${filename}.xlsx`);
    },
    
    async GetEmbarazada() {
      try {
        this.embarazadas = await EmbarazadaService.GetEmbarazada();
      } catch {
        DialogosService.mostrarNotificacionError(
          "No se puede obtener la lista de embarazada"
        );
      }
    },
  },
};
</script>
<style>
.celdanm {
  width: min-content;
  background: #4e7d96 !important;
  color: white !important;
}
.nombretb {
  background: #4e7d96 !important;
  color: white !important;
}
.cell-class {
  width: auto;
}
</style>
