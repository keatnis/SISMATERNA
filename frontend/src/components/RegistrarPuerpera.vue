<template>
  <!-- <div class="columns"> -->
  <b-steps v-model="activeStep" :animated="isAnimated" :rounded="isRounded" :has-navigation="hasNavigation"
    :icon-prev="prevIcon" :icon-next="nextIcon" :mobile-mode="mobileMode">
    <div class="has-text-centered">
      <article class="panel is-info centered">
        <div class="panel-heading">Registrar puerpera</div>
      </article>
    </div>
    <div>
      <div class="columns">
        <div class="column is-half">
          <b-field label="No. Progreso">
            <b-input disabled="»disabled»" v-model="detalles.id" placeholder=""></b-input>
          </b-field>
          <b-field label="Nombre de puerpera">
            <b-input v-model="detalles.nombrepuerpera" placeholder="Nombre completo" disabled="»disabled»"></b-input>
          </b-field>

          <b-field label=" Consulta de 7 días">
            <b-datepicker v-model="detalles.consultasiete" placeholder="Fecha consulta de 7 días" icon="calendar-today"
              :locale="locale" editable>
            </b-datepicker>
          </b-field>

          <b-field label="Consulta de 40 días">
            <b-datepicker v-model="detalles.consultacuarenta" placeholder="Fecha consulta de 40 días"
              icon="calendar-today" :locale="locale" editable>
            </b-datepicker>
          </b-field>

          <b-field label="Atención de parto">
            <b-select v-model="detalles.atencionparto" placeholder="¿Quién atendió el parto?" expanded>
              <option value="Médico">Médico</option>
              <option value="Enfermeraa">Enfermera</option>
              <option value="Partera Trad.">Partera Trad.</option>
              <option value="Partera Prof.">Partera Prof.</option>
              <option value="TAPS">TAPS</option>
              <option value="Familiar">Familiar</option>
              <option value="Ninguno">Ninguno</option>
            </b-select>
          </b-field>
          <b-field label="Resolución del embarazo">
            <b-select v-model="detalles.resolucion" placeholder="Resolución del embarazo" expanded>
              <option value="Vaginal">Vaginal</option>
              <option value="Cesarea ">Cesarea</option>
              <option value="Legrado (aborto)">Legrado (aborto)</option>
              <option value="Legrado (óbito)">Legrado (óbito)</option>
              <option value="Ninguno">Ninguno</option>
            </b-select>
          </b-field>

          <b-field label="APEO">
            <b-select placeholder="APEO" v-model="detalles.apeo" expanded>
              <option value="DIU">DIU</option>
              <option value="OTB">OTB</option>
              <option value="Implante">Implante</option>
              <option value="Ninguno">Ninguno</option>
            </b-select>
          </b-field>
          <b-field label="Consulta pregestacional"></b-field>
          <b-radio v-model="detalles.pregestacional" :native-value="1" >Si</b-radio>
          <b-radio v-model="detalles.pregestacional" :native-value="0" >No</b-radio>
          <b-field label="Derechohabiencia">
            <b-select placeholder="Derechohabiencia" v-model="detalles.derechohabiencia" disabled expanded>
              <option value="INSABI">INSABI</option>
              <option value="IMSS">IMSS</option>
              <option value="ISSSTE">ISSSTE</option>
              <option value="SEDENA">SEDENA</option>
              <option value="PEMEX">PEMEX</option>
              <option value="MARINA">MARINA</option>
              <option value="Otro">Otro</option>
            </b-select>
          </b-field>
        </div>
        <div class="column is-half">
          <b-field label="CURP">
            <b-input v-model="detalles.descripcion" placeholder="Clave Unica de Registro de Población" icon="magnify">
            </b-input>
          </b-field>
          <b-field label="Edad">
            <b-input disabled="»disabled»" v-model="detalles.edad"></b-input>
          </b-field>

          <b-field label="Consulta de 28 días">
            <b-datepicker v-model="detalles.consultaveinte" placeholder="Fecha consulta de 28 días" icon="calendar-today"
              :locale="locale" editable>
            </b-datepicker>
          </b-field>
          <b-field label="Signos de alarma">
            <b-select v-model="detalles.signos" placeholder="Signos" expanded>
              <option value="Fiebre">Fiebre</option>
              <option value="Dolor Pélvico">Dolor Pélvico</option>
              <option value="Flujo Vaginal">Flujo Vaginal</option>
              <option value="Fiebre/Dolor Pelvico">Fiebre/Dolor Pelvico</option>
              <option value="Fiebre/Flujo Vaginal">Fiebre/Flujo Vaginal</option>
              <option value=">Dolor Pelvico/Flujo Vaginalo">Dolor Pelvico/Flujo Vaginal</option>
              <option value="Fiebre/Dolor Pelvico/Flujo Vaginal">
                Fiebre/Dolor Pelvico/Flujo Vaginal
              </option>
              <option value="Cefalea">Cefalea</option>
              <option value="Ninguno">Ninguno</option>
            </b-select>
          </b-field>

          <b-field label="Lugar del parto">
            <b-select v-model="detalles.lugarparto" placeholder="¿Dónde atendió el parto?" expanded>
              <option value="H.G. Tlapa">H.G. Tlapa</option>
              <option value="HMNIG">HMNIG</option>
              <option value="H.C. Acatepec">H.C. Acatepec</option>
              <option value="H.C. Alcozauca">H.C. Alcozauca</option>
              <option value="H.C. Huamuxtitlán">H.C. Huamuxtitlán</option>
              <option value="H.C. Malinaltepec">H.C. Malinaltepec</option>
              <option value="H.C. Olinalá">H.C. Olinalá</option>
              <option value="H.C. Xochihuehuetlán">H.C. Xochihuehuetlán</option>
              <option value="H.C. Zapotitlán Tablas">H.C. Zapotitlán Tablas</option>
              <option value="Centro de Salud">Centro de Salud</option>
              <option value="Hogar">Hogar</option>
              <option value="Otro hospital…">Otro hospital…</option>
            </b-select>
          </b-field>

          <b-field label="No. de producto">
            <b-select v-model="detalles.producto" placeholder="No. de producto" expanded>
              <option value="Único">Único</option>
              <option value="Gemelar">Gemelar</option>
              <option value="Trillizos">Trillizos</option>
              <option value="Cuatrillizos">Cuatrillizos</option>
            </b-select>
          </b-field>

          <b-field label="Puerpera aceptante">
            <b-select v-model="detalles.aceptante" placeholder="Puerpera acentante" expanded>
              <option value="DIU">DIU</option>
              <option value="OTB">OTB</option>
              <option value="Implante">Implante</option>
              <option value="H.I. Bimensual">H.I. Bimensual</option>
              <option value="Ninguno">Ninguno</option>
            </b-select>
          </b-field>

          <b-field label="Fecha de parto">
            <b-datepicker placeholder="Fecha del evento" icon="calendar-today" :locale="locale" disabled editable>
            </b-datepicker>
          </b-field>

          <b-field>
            <b-button @click="guardar()" type="is-success">Guardar</b-button>
          </b-field>
        </div>
      </div>
    </div>
  </b-steps>
</template>
<script>
import Utiles from "../services/Utiles";
import DialogosService from "../services/DialogosService";
import PuerperaService from "../services/PuerperaService";
import { DialogProgrammatic as Dialog } from "buefy";
export default {
  data: () => ({
    selected: null,
    edad: null,
    activeStep: 0,

    isProfileSuccess: false,
    checkbox: false,
    mobileMode: "minimalist",
    hasNavigation: true,
    locale: "es-MX",
    isAnimated: true,
    isRounded: true,
    isStepsClickable: false,
    prevIcon: "chevron-left",
    nextIcon: "chevron-right",
    showSocial: false,
    checkboxGroup: [],
    detalles: {
      consultasiete: null,
      consultaveinte: null,
      consultacuarenta: null,
      signos: "",
      atencionparto: "",
      lugarparto: "",
      resolucion: "",
      producto: "",
      apeo: "",
      aceptante: "",
      pregestacional: 0,
    },
  }),
  computed: {},

  methods: {
    formatearFecha(fecha) {
      return Utiles.obtenerCadenaFecha(fecha);
    },

    confirm() {
      Dialog.confirm({
        title: "Guardar registro.",
        message: "Confirmar para guardar datos a la base de datos.",
        cancelText: "Cancelar",
        confirmText: "Aceptar",
        type: "is-success",
        onConfirm: () => this.guardar(),
      });
    },
    async guardar() {
      if (!this.detalles.consultasiete == null) {
        return DialogosService.mostrarNotificacionError("Campos vacios");
      }
      const cargaUtil = {
        consultasiete: Utiles.obtenerCadenaFecha(this.detalles.consultasiete),
        consultaveinte: Utiles.obtenerCadenaFecha(this.detalles.consultaveinte),
        consultacuarenta: Utiles.obtenerCadenaFecha(
          this.detalles.consultacuarenta
        ),
        signos: this.detalles.signos,
        atencionparto: this.detalles.atencionparto,
        lugarparto: this.detalles.lugarparto,
        resolucion: this.detalles.resolucion,
        producto: this.detalles.producto,
        apeo: this.detalles.apeo,
        aceptante: this.detalles.aceptante,
        pregestacional: this.detalles.pregestacional,
      };

      const respuesta = await PuerperaService.insertPuerpera(cargaUtil);
      if (respuesta != true) {
        console.log("respuesta save", respuesta);
        DialogosService.mostrarError("registro con errores");
      } else {
        DialogosService.mostrarExitoso(
          "Datos guardados correctament en la base de datos"
        );
        this.detalles = {
          consultasiete: null,
          consultaveinte: null,
          consultacuarenta: null,
          signos: "",
          atencionparto: "",
          lugarparto: "",
          resolucion: "",
          lenguaIndigena: "",
          producto: "",
          apeo: "",
          aceptante: "",
          pregestacional: false,
        };
      }
    },
  },
};
</script>
