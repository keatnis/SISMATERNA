<template>
  
    <section>
      <article class="panel is-info">
    <p class="panel-heading">
      Bucar embarazada:
    </p></article>
    
    

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
      <b-table :data="embarazadas" :columns="columns" :sticky-header="stickyHeaders" ></b-table>  
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
          field: "nuevoingreso",
          label: "Nuevo ingreso",

        },

        {
          field: "noExpediente",
          label: "Numero de expediente clinico",
          width: "70",
         
        },
        {
          field: "nombre",
          label: "Nombre de la embarazada",
          numeric: true,
          sticky: true,
          
          centered: true,

          
        },

        {
          field: "FechaNacimiento",
          label: "Fecha de nacimiento ",
          centered: true,

        },

        {
          field: "edad",
          label: "Edad",

        },
        {
          field: "domicilioReferencia",
          label: "Direccion o referencia ",
          centered: true,

        },

        {
          field: "localidad",
          label: "Localidad",

        },
        {
          field: "municipio",
          label: "Municipio",

        },
        {
          field: "telefono",
          label: "Numero telefonico",
          centered: true,

        },
        {
          field: "lenguaIndigena",
          label: "Habla lengua indigena",
          centered: true

        },
        {
          field: "emigro",
          label: "Emigro"

        },
        {
          field: "derechohabiencia",
          label: "Derechohabiencia"
        },
        {
          field: "violencia",
          label: "Violencia"
        },
        {
          field: "detecciones",
          label: "Detecciones"
        },

        {
          field: "ConsultaPregestacional",
          label: "Consulta de riesgo pregestacional",
          centered:true

        },

        {
          field: "comorbilidades",
          label: "Comorbilidades",
          centered: true
        },


        {
          field: "gestas",
          label: "Gestas",

        },
        {
          field: "paras",
          label: "Paras",
          centered: true,

        },
        {
          field: "abortos",
          label: "Abortos",
          centered: true,
          
        },
        {
          field: "cesareas",
          label: "Cesareas",
          centered: true,

        },

        {
          field: "FechaUltimoEvento",
          label: "Fecha ultimo parto, cesarea o aborto",
          centered: true,

        },
        {
          field: "presentComplicaciones",
          label: "¿Presentó complicaciones en el embarazo previo?",
          centered: true,

        },
        {
          field: "complicaciones",
          label: "¿Que complicaciones presento?",
          centered: true
        },
        {
          field: "FechaUlmaMenstruacion",
          label: "Fecha ultima menstruación",
          centered: true,
          
        },
        {
          field: "FechaProbableParto",
          label: "Fecha probable de parto",
          centered: true,
          
        },
        {
          field: "trimestre",
          label: "Trimestre actual",
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
        {
          field: "noConsultasEmbarazo",
          label: "No.de consulta durante el embarazo",
          centered: true,
          
        },
        {
          field: "noConsultasMes",
          label: "No. consulta en el mes",
          centered: true,
          
        },
        {
          field: "abortos",
          label: "Abortos",
          centered: true,
          
        },
        {
          field: "rubeola",
          label: "Rubeola",
          centered: true,
          
        },
        {
          field: "FechaVacunaTDPrimera",
          label: "Fecha TD primera",
          centered: true,
          
        },
        {
          field: "FechaVacunaTDSegunda",
          label: "Fecha TD segunda",
          centered: true,
          
        },
        {
          field: "FechaVacunaTDRefuerzo",
          label: "Fecha TD refuerzo",
          centered: true,
          
        },
        {
          field: "FechaVacunaTDPA",
          label: "Fecha TDPA",
          centered: true,
          
        },
        {
          field: "FechaVacunaInfluenza",
          label: "Fecha influenza estacional ",
          centered: true,
          
        },
        {
          field: "vacunaCOVID",
          label: "AntiCovid-19",
          centered: true,
          
        },
        {
          field: "grupoRH",
          label: "Grupo y RH",
          centered: true,
          
        },
        {
          field: "ego",
          label: "EGO",
          centered: true,
          
        },
        {
          field: "BiometriaHematica",
          label: "Biometria ematica",
          centered: true,
          
        },
        {
          field: "leucocitos",
          label: "Leucocitos",
          centered: true,
          
        },
        {
          field: "plaquetas",
          label: "Plaquetas",
          centered: true,
          
        },
        {
          field: "vdrl",
          label: "VDRL",
          centered: true,
          
        },
        {
          field: "vih",
          label: "VIH",
          centered: true,
          
        },
        {
          field: "estadoGlucosa",
          label: "estadoGlucosa",
          centered: true,
          
        },
        {
          field: "resultadoGlucosa",
          label: "resultadoGlucosa",
          centered: true,
          
        },
        {
          field: "caracteristicasFetls",
          label: "Caracteristicas fetales",
          centered: true,
          
        },
        {
          field: "malformaciones",
          label: "Malformaciones",
          centered: true,
          
        },
        {
          field: "LiquidoAmiotico",
          label: "Liquido amiotico",
          centered: true,
          
        },
        {
          field: "placenta",
          label: "Placenta",
          centered: true,
          
        },
        {
          field: "FechaProbableUSG",
          label: "Fecha Probable de parto por USG",
          centered: true,
          
        },
        {
          field: "referencia",
          label: "Referencia",
          centered: true,
          
        },
        {
          field: "acudio_refe",
          label: "Acudio a refenrecia",
          centered: true,
          
        },
        {
          field: "contrareferencia",
          label: "Contrareferencia recibida",
          centered: true,
          
        },
        {
          field: "planSeguridad",
          label: "¿Se le realizo plan de seguridad?",
          centered: true,
          
        },
        {
          field: "lugarParto",
          label: "¿Donde atendera parto?",
          centered: true,
          
        },
        {
          field: "quienAtenderaParto",
          label: "¿Quien atendera parto?",
          centered: true,
          
        },
        {
          field: "transporteAME",
          label: "Cuenta con transporte AME",
          centered: true,
          
        },
        {
          field: "FechaEvento",
          label: "Fecha de vento",
          centered: true,
          
        },
        {
          field: "observaciones",
          label: "Observaciones",
          centered: true,
          
        },
        {
          field: "diagnostico",
          label: "Diagnostico",
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
.is-sticky-column-one {
  background: #7cc4eb !important;
  color: white !important;
}

.is-sticky-column-two {
  background: #6ca8ee !important;
  color: white !important;
}
</style>