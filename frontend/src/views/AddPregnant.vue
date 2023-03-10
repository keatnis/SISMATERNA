<template>
  <div>
    <hero-bar> Registro Embarazada </hero-bar>
    <section class="section is-main-section">
      <card-component title="Datos de Paciente" icon="ballot">
        <form>
          <div class="columns">
            <div class="column is-one-third">
              <b-field label="Nombre Completo">
                <b-input
                  v-model="detalles.NombreEmbarazada"
                  placeholder="Nombre completo de paciente"
                />
              </b-field>
              <b-field label="Apellido Materno">
                <b-input
                  v-model="detalles.NombreEmbarazada"
                  placeholder="Nombre completo de paciente"
                />
              </b-field>
              <b-field label="CURP">
                <b-input
                  v-model="detalles.curp"
                  placeholder="Clave Unica de Registro de Población"
                  maxlength="18"
                  required
                />
              </b-field>
              <b-field label="No. Expediente clínico">
                <b-input
                  v-model="detalles.noExpediente"
                  placeholder="No. Expediente"
                  type="number"
                  required
                />
              </b-field>
              <b-field label="Fecha de nacimiento">
                <b-datepicker
                  v-model="detalles.fechaNacimiento"
                  placeholder="Fecha de nacimiento"
                  icon="calendar-today"
                  :datetime-formatter="formatearFecha"
                  :locale="locale"
                  editable
                  @input="calcularEdad()"
                />
                Edad: {{ detalles.edad }}
              </b-field>
              <b-field label="Lengua Indígena">
                <b-select
                  v-model="detalles.lenguaIndigena"
                  placeholder="Lengua Indígena"
                  expanded
                >
                  <option
                    v-for="(lenguaIndigena, index) in datajs.lenguas"
                    :key="index"
                    :value="lenguaIndigena"
                  >
                    {{ lenguaIndigena }}
                  </option>
                </b-select>
              </b-field>
              <b-field label="Municipio">
                <b-select
                  v-model="municipioSeleccionado"
                  placeholder="Seleccionar municipio"
                  expanded
                >
                  <option
                    v-for="option in municipios"
                    :key="option.id"
                    :value="{
                      id: option.id_municipio,
                      text: option.nombreMunicipio,
                    }"
                  >
                    {{ option.nombreMunicipio }}
                  </option>
                </b-select>
              </b-field>
              <b-field label="Localidad">
                <b-select
                  v-model="localidadSeleccionada"
                  placeholder="Seleccionar localidad"
                  expanded
                >
                  <option
                    v-for="option in localidades"
                    :key="option.id"
                    :value="{
                      id: option.id_localidad,
                      text: option.NombreLocalidad,
                    }"
                  >
                    {{ option.NombreLocalidad }}
                  </option>
                </b-select>
              </b-field>
              <b-field label="Emigro" horizontal>
                <b-radio v-model="detalles.emigro" :native-value="1">
                  Si
                </b-radio>
                <b-radio v-model="detalles.emigro" :native-value="0">
                  No
                </b-radio>
              </b-field>

              <b-field label="Dirección">
                <b-input
                  v-model="detalles.Direccion"
                  placeholder="Dirección con referencia"
                />
              </b-field>
              <b-field label="Número teléfono">
                <b-input
                  v-model="detalles.TelefonoEmbarazada"
                  placeholder="Teléfono fijo"
                />
              </b-field>
              <b-field label="Derechohabiencia">
                <b-select v-model="detalles.Derechohabiencia" expanded>
                  <option
                    v-for="(
                      derechohabiencia, index
                    ) in datajs.derechohabiencias"
                    :key="index"
                    :value="derechohabiencia"
                  >
                    {{ derechohabiencia }}
                  </option>
                </b-select>
              </b-field>
              <div class="field">
                <b-field label="Detenciones">
                  <checkbox-radio-picker
                    v-model="detalles.detenciones"
                    :options="{
                      lorem: 'Exploración clínica de mama normal',
                      explo: 'Exploración clínica de mama alterado',
                      Papa_normal: 'Papanicolau normal',
                      papa_lesiones: 'Papanicolau con lesiones',
                      papa_cancer: 'Papanicolau cancer IN SITU',
                      papa_invasor: 'Papanicolau invasor',
                      no: 'No se realizo',
                    }"
                  />
                </b-field>
                <p class="content">
                  <b>Selección:</b>
                  {{ detalles.detenciones }}
                </p>
              </div>
            </div>

            <div class="column is-one-third">
              <b-field label="¿Asistió a consulta pregestacional?"></b-field>
              <b-radio v-model="detalles.AsistenciaPreg" :native-value="1"
                >Si</b-radio
              >
              <b-radio v-model="detalles.AsistenciaPreg" :native-value="0"
                >No</b-radio
              >
              <b-datepicker
                :v-model="
                  detalles.AsistenciaPreg
                    ? detalles.consultaPregestacional
                    : detalles.citaConsultaPregestional
                "
                :placeholder="
                  detalles.AsistenciaPreg
                    ? 'Fecha consulta'
                    : 'Fecha de próxima cita'
                "
                icon="calendar-today"
                :datetime-formatter="formatearFecha"
                :locale="locale"
                editable
              >
              </b-datepicker>
              <hr />
              <b-field label="Comorbilidades">
                <section>
                  <CheckboxRadioPicker
                    v-model="detalles.comorbilidades"
                    class="container"
                    :options="{
                      sobrepeso: 'Sobrepeso',
                      obesidad: 'Obesidad',
                      diabetis_mellitus: 'Diabetes Mellitus',
                      cardiopatia: 'Cardiopatia',
                      epilepsia: 'Epilepsia',
                      sifilis: 'Sifilis',
                      vih: 'VIH',
                      hepatitis: 'Hepatitis',
                      lupus: 'Lupus',
                      cancer: 'Cancer',
                      tuberculosis: 'Tuberculosis',
                      covid: 'COVID-19',
                      anemia: 'Anemia',
                      cervi: 'Cervicovaginitis',
                      artritis_reumatoide: 'Artritis reumatoide',
                      hipertension_arterial: 'Hipertensión arterial',
                      infeccion_urinaria: 'Infección de Vias Urinarias',
                      ninguno: 'Ninguno',
                    }"
                  />

                  <p class="content">
                    <b>Selección:</b>
                    {{ detalles.comorbilidades }}
                  </p>
                </section>
              </b-field>
              <hr />
              <b-field label="Qué metodo Planificación Familiar usa?">
                <b-select v-model="detalles.metodopf" expanded>
                  <option
                    v-for="(metodopf, index) in datajs.metodospf"
                    :key="index"
                    :value="metodopf"
                  >
                    {{ metodopf }}
                  </option>
                </b-select>
              </b-field>
              <h1 class="subtitle">Antecedentes Gineco obstétricos</h1>
              <b-field label="Gestas">
                <b-numberinput
                  v-model="detalles.Gestas"
                  type="number"
                  placeholder="Número de gestas"
                />
              </b-field>
              <b-field label="Cesareas">
                <b-numberinput
                  v-model="detalles.Cesareas"
                  type="number"
                  placeholder="Número de cesareas"
                />
              </b-field>
              <b-field label="Paras">
                <b-numberinput
                  v-model="detalles.Paras"
                  type="number"
                  placeholder="Numero de paras"
                />
              </b-field>
              <b-field label="Abortos">
                <b-numberinput
                  v-model="detalles.Abortos"
                  type="number"
                  placeholder="Numero de abortos"
                />
              </b-field>
              <hr />
              <b-field
                label="¿Presentó complicaciones en el embarazo previo?"
              />
              <b-radio
                v-model="detalles.presentComplicaciones"
                :native-value="1"
              >
                Si
              </b-radio>
              <b-radio
                v-model="detalles.presentComplicaciones"
                :native-value="0"
              >
                No
              </b-radio>
              <b-select
                :v-model="detalles.presentComplicaciones"
                placeholder="¿Qué complicaciones presento?"
                :disabled="detalles.presentComplicaciones ? false : true"
                expanded
              >
                <option
                  v-for="(complicacion, index) in datajs.complicaciones"
                  :key="index"
                  :value="detalles.presentComplicaciones ? complicacion : null"
                >
                  {{ complicacion }}
                </option>
              </b-select>
            </div>
            <div class="column is-one-third">
              <h1 class="subtitle">Embarazo Actual</h1>
              <b-field label="Fecha de ultima menstruación">
                <b-datepicker
                  v-model="detalles.fechaUlmaMenstruacion"
                  placeholder="Fecha de ultima menstruación"
                  icon="calendar-today"
                  :datetime-formatter="formatearFecha"
                  :locale="locale"
                  editable
                >
                </b-datepicker>
              </b-field>
              <b-field label="Trimestre de Capacitación"></b-field>
              <checkbox-radio-picker
                v-model="trimestre_capacitacion"
                :options="{
                  1: 'Primer Trimestre',
                  2: 'Segundo Trimestre',
                  3: 'Tercer Trimestre',
                }"
              ></checkbox-radio-picker>
              <b-field label="Fecha de ultimo evento( Parto, cesarea, aborto)">
                <b-datepicker
                  v-model="detalles.fechaUltimoEvento"
                  placeholder="Fecha de ultimo evento"
                  icon="calendar-today"
                  :datetime-formatter="formatearFecha"
                  :locale="locale"
                  editable
                />
              </b-field>

              <b-field label="Semana de gestación actual (SGA)">
                <b-input
                  v-model="detalles.SGA"
                  placeholder="Semana de gestación actual"
                  type="is-warning"
                >
                </b-input>
              </b-field>
              <b-field label="Fecha probable de parto (FPP)">
                <b-datepicker
                  v-model="detalles.fechaProbableParto"
                  placeholder="FPP"
                  icon="calendar-today"
                  :datetime-formatter="formatearFecha"
                  :locale="locale"
                  editable
                >
                </b-datepicker>
              </b-field>
              <b-field label="No. Consultas"></b-field>
              <b-select :v-model="detalles.noConsultasEmbarazo" expanded>
                <option
                  v-for="(
                    noConsultasEmbarazo, index
                  ) in datajs.consultasEmbarazo"
                  :key="index"
                  :value="noConsultasEmbarazo"
                >
                  {{ noConsultasEmbarazo }}
                </option>
              </b-select>
              <b-field label="Fecha de consulta">
                <b-datepicker
                  v-model="detalles.fechaConsulta"
                  placeholder="Fecha de consulta "
                  icon="calendar-today"
                  :datetime-formatter="formatearFecha"
                  :locale="locale"
                  editable
                >
                </b-datepicker>
              </b-field>
              <b-field label="No. Total de consulta en el mes">
                <b-numberinput
                  disabled
                  v-model="detalles.noConsultasMes"
                  placeholder="No. Total de consulta en el mes "
                >
                </b-numberinput>
              </b-field>
            </div>
          </div>
        </form>
      </card-component>
    </section>
    <section class="section is-main-section">
      <card-component title="Esquema de vacunación y obstetrico " icon="needle">
        <div class="columns">
          <div class="column is-one-third">
            <b-field label="Rubeola:">
              <b-radio v-model="detalles.rabiola" :native-value="1"
                >Aplicada previo al embarazo</b-radio
              >
              <b-radio v-model="detalles.rabiola" :native-value="0"
                >No cuenta con vacuna</b-radio
              >
            </b-field>
            <hr />
            <h1 class="subtitle">TD</h1>
            <b-field label="Primera:" horizontal>
              <b-datepicker
                v-model="detalles.fechaVacunaTDPrimera"
                placeholder="Fecha TD Primera"
                icon="calendar-today"
                :datetime-formatter="formatearFecha"
                :locale="locale"
                editable
              ></b-datepicker>
            </b-field>
            <b-field label="Segunda:" horizontal>
              <b-datepicker
                v-model="detalles.fechaVacunaTDSegunda"
                placeholder="Fecha TD Segunda"
                icon="calendar-today"
                :datetime-formatter="formatearFecha"
                :locale="locale"
                editable
              ></b-datepicker>
            </b-field>

            <b-field label="Refuerzo" horizontal>
              <b-datepicker
                v-model="detalles.fechaVacunaTDRefuerzo"
                placeholder="Fecha de vacunación TD refuerzo"
                icon="calendar-today"
                :datetime-formatter="formatearFecha"
                :locale="locale"
                editable
              >
              </b-datepicker>
            </b-field>
            <hr />
            <h1 class="subtitle">TDPA:</h1>
            <b-field label="Refuerzo:" horizontal>
              <b-datepicker
                v-model="detalles.fechaVacunaTDPA"
                placeholder="Fecha de vacunación TDPA"
                icon="calendar-today"
                :datetime-formatter="formatearFecha"
                :locale="locale"
                editable
              ></b-datepicker>
            </b-field>
            <hr />
            <b-field label="Fecha de vacunación Influenza estacional:">
              <b-datepicker
                v-model="detalles.fechaVacunaInfluenza"
                placeholder="Fecha de vacunación Influenza estacional "
                icon="calendar-today"
                :datetime-formatter="formatearFecha"
                :locale="locale"
                editable
              >
              </b-datepicker>
            </b-field>

            <b-field label="Vacuna ANTICOVID-19:">
              <b-radio v-model="detalles.vacunaCOVID" :native-value="1"
                >Se aplicó</b-radio
              >
              <b-radio v-model="detalles.vacunaCOVID" :native-value="0"
                >No se aplico</b-radio
              >
            </b-field>
          </div>
          <div class="column is-one-third">
            <h1 class="subtitle">Laboratorio y Gabinete</h1>
            <b-field label="GRUPO Y RH">
              <b-select
                v-model="detalles.grupoRH"
                placeholder="Seleccione el grupo y RH"
                expanded
              >
                <option
                  v-for="(grupo, index) in datajs.grupoRH"
                  :key="index"
                  :value="grupo"
                >
                  {{ grupo }}
                </option>
              </b-select>
            </b-field>
            <b-field label="EGO Y/O Tiras Urianalisis"> </b-field>
            <b-select v-model="detalles.ego" placeholder="EGO" expanded>
              <option
                v-for="(eg, index) in datajs.ego"
                :key="index"
                :value="eg"
              >
                {{ eg }}
              </option>
            </b-select>

            <b-field label="Biometria ematica">
              <b-select
                v-model="detalles.BiometriaHematica"
                placeholder="Biometria"
                expanded
              >
                <option
                  v-for="(biometria, index) in datajs.BiometriaHematica"
                  :key="index"
                  :value="biometria"
                >
                  {{ biometria }}
                </option>
              </b-select>
            </b-field>
            <b-field label="VDRL">
              <b-select v-model="detalles.vdrl" placeholder="VDRL" expanded>
                <option
                  v-for="(vdr, index) in datajs.VDRL"
                  :key="index"
                  :value="vdr"
                >
                  {{ vdr }}
                </option>
              </b-select>
            </b-field>
            <b-field label="VIH">
              <b-select v-model="detalles.vih" placeholder="VIH" expanded>
                <option value="Reactivo VIH">Reactivo</option>
                <option value="NoReactivo VIH">No reactivo</option>
                <option value="NoRealizo VIH">No se realizó</option>
              </b-select>
            </b-field>

            <b-field label="Tiras Glucosa:" />
            <b-field label="Estado de Glucosa">
              <b-select
                v-model="detalles.estadoGlucosa"
                placeholder="Biometria"
                expanded
              >
                <option
                  v-for="(estado, index) in datajs.estadoGlucosa"
                  :key="index"
                  :value="estado"
                >
                  {{ estado }}
                </option>
              </b-select>
            </b-field>
            <b-field label="Resultado Glucosa">
              <b-select
                v-model="detalles.resultadoGlucosa"
                placeholder="Biometria"
                expanded
              >
                <option
                  v-for="(resultado, index) in datajs.resultadoGlucosa"
                  :key="index"
                  :value="resultado"
                >
                  {{ resultado }}
                </option>
              </b-select>
            </b-field>
          </div>
          <div class="column is-one-third">
            <h1 class="subtitle">ULTRASONIDO</h1>
            <b-field label="Caracteristicas fetales:" position="is-left" >
              <b-select
                placeholder="Caract. Fetales "
                v-model="detalles.caracteristicasFetls"
                expanded
              >
                <option
                  v-for="(caract, index) in datajs.caracteristicasFetls"
                  :key="index"
                  :value="caract"
                >
                  {{ caract }}
                </option>
              </b-select>
            </b-field>
            <b-field label="Malformaciones congenitas:"  position="is-left" >
              <b-select
                placeholder="Malformaciones"
                v-model="detalles.malformaciones"
                expanded
              >
                <option
                  v-for="(malf, index) in datajs.malformaciones"
                  :key="index"
                  :value="malf"
                >
                  {{ malf }}
                </option>
              </b-select>
            </b-field>

            <b-field label="Liquido Amniótico:"  position="is-left" >
              <b-select
                placeholder="Liquido Amniótico"
                v-model="detalles.liquidoAmiotico"
                expanded
                 >
                <option
                  v-for="(amio, index) in datajs.liquidoAmiotico"
                  :key="index"
                  :value="amio"
                >
                  {{ amio }}
                </option>
              </b-select>
            </b-field>
            <b-field label="Placenta:"  position="is-left" >
              <b-select
                placeholder="Placenta"
                v-model="detalles.placenta"
                expanded
              >
                <option
                  v-for="(placent, index) in datajs.placenta"
                  :key="index"
                  :value="placent"
                >
                  {{ placent }}
                </option>
              </b-select>
            </b-field>
            <b-field label="Fecha probable parto por USG (1er Trimestre)">
              <b-datepicker
                v-model="detalles.fechaProbableUSG"
                placeholder="Fecha probable de parto por USG (1Er Trimestre)"
                icon="calendar-today"
                :datetime-formatter="formatearFecha"
                :locale="locale"
                editable
              >
              </b-datepicker>
            </b-field>
            <hr/>
            <H1 class="subtitle">Trabajo social</H1>  
            <b-field label="Referencia"/>
            <b-radio v-model="detalles.referencia" :native-value="1"
                >Si</b-radio
              >
              <b-radio v-model="detalles.referencia" :native-value="0"
                >No</b-radio
              >
              <b-field label="¿Acudió referencia?"></b-field>
            <b-radio v-model="detalles.acudio_refe" :native-value="1"
                >Si</b-radio
              >
              <b-radio v-model="detalles.acudio_refe" :native-value="0"
                >No</b-radio
              >
          </div>
        </div>
      </card-component>
    </section>
    <section >
      <Card-component title="Componente comunitario y evento obtétrico">
        <form>
          <b-field label="Se le realizó el plan de seguridad">
            <b-radio v-model="detalles.planSeguridad" :native-value="1"
                >Si</b-radio
              >
              <b-radio v-model="detalles.planSeguridad" :native-value="0"
                >No</b-radio
              >
              </b-field>

        </form>
      </Card-component>
    </section>
    <section>
    <hr/>
      <CardComponent title="Otros">
        
        <b-field horizontal label="Diagnostico">
            <b-input type="textarea" placeholder="Establecer riesgo y motivo de riesgo"></b-input>
        </b-field>
        <b-field horizontal label="Observaciones">
            <b-input type="textarea" placeholder="Ej. Menor de edad, multiparidad, etc."></b-input>
        </b-field>
      </CardComponent>
    </section>
  </div>
</template>

<script>
import HeroBar from "@/components/HeroBar.vue";
import CardComponent from "@/components/CardComponent.vue";
import CheckboxRadioPicker from "@/components/CheckboxRadioPicker.vue";
import datajs from "@/store/arrayRegistroEmb.js";

import Utiles from "@/services/Utiles";
import DialogosService from "@/services/DialogosService";
import EmbarazadaService from "@/services/EmbarazadaService";
import { DialogProgrammatic as Dialog } from "buefy";


export default {
  name: "AddPregnant",
  components: {
    HeroBar,
    CardComponent,
    CheckboxRadioPicker
   
  },
  data: () => ({
    datajs,
    selected: null,
    edad: null,
    trimestre_capacitacion: [],
    municipios: [],
    localidades: [],
    localidadSeleccionada: "",
    municipioSeleccionado: "",
    formatearFecha: Utiles.formatearFechaSegunLocale,
    fechaInicio: null,
    checkbox: false,
    mobileMode: "minimalist",
    locale: "es-MX",
    checkboxGroup: [],
    detalles: {
      detenciones: [],
      noExpediente: "",
      NombreEmbarazada: "",
      curp: "",
      Direccion: "",
      Derechohabiencia: "",
      TelefonoEmbarazada: null,
      DondeMigro: 0,
      fechaNacimiento: null,
      Gestas: 0,
      Paras: 0,
      Abortos: 0,
      Cesareas: 0,
      consultaPregestacional: null,
      fechaUltimoEvento: null,
      fechaUlmaMenstruacion: null,
      fechaProbableParto: null,
      fechaConsulta: null,
      fechaVacunaTDPrimera: null,
      fechaVacunaTDSegunda: null,
      fechaVacunaTDRefuerzo: null,
      fechaVacunaTDPA: null,
      fechaVacunaInfluenza: null,
      fechaProbableUSG: null,
      fechaEvento: null,
      edad: "",
      lenguaIndigena: "",
      emigro: 0,
      violencia: "",
      comorbilidades: [],
      AsistenciaPreg: 0,
      presentComplicaciones: 0,
      complicaciones: "",
      SGA: "",
      noConsultasEmbarazo: 0,
      noConsultasMes: 1,
      rubeola: 0,
      vacunaCOVID: 0,
      grupoRH: "",
      ego: "",
      BiometriaHematica: "",
      leucocitos: "",
      plaquetas: "",
      vdrl: "",
      vih: "",
      estadoGlucosa: "",
      resultadoGlucosa: "",
      caracteristicasFetls: "",
      malformaciones: "",
      liquidoAmiotico: "",
      placenta: "",
      referencia: "",
      acudio_refe: "",
      contrareferencia: "",
      planSeguridad: "",
      lugarParto: "",
      quienAtenderaParto: "",
      transporteAME: "",
      observaciones: "",
      diagnostico: "",
      metodopf: "",
      citaConsultaPregestional: "",
    },
  }),
  computed: {
    /*
                      selectedString() {
                      return this.selected ? this.selected.toDateString() : "";
                    },
                    */
  },
  async mounted() {
    await this.obtenerMunicipio();
  },
  methods: {
    calcularEdad() {
      var hoy = new Date();
      //    const fechaCumple = Utiles.obtenerCadenaFecha( this.detalles.fechaNacimiento );
      var cumpleanos = new Date(this.detalles.fechaNacimiento);
      var edad = hoy.getFullYear() - cumpleanos.getFullYear();
      var mes = hoy.getMonth() - cumpleanos.getMonth();
      // ()  var dia = cumpleanos.getDay();

      if (mes < 0 || (mes === 0 && hoy.getDate() < cumpleanos.getDate())) {
        edad--;
      }
      // this.detalles.fechaNacimiento = cumpleanos;
      this.detalles.edad = edad;
      return edad;
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
    async obtenerMunicipio() {
      try {
        this.municipios = await EmbarazadaService.obtenerMunicipios();
      } catch (e) {
        DialogosService.mostrarNotificacionError(
          "No se pudo obtener la lista de municipios..."
        );
      }
    },

    async obtenerLocalidades() {
      try {
        if (this.municipioSeleccionado.id === "undefined") {
          return;
        } else {
          this.localidades = await EmbarazadaService.obteneLocalidadesById(
            this.municipioSeleccionado.id
          );
          return;
        }
      } catch (e) {
        DialogosService.mostrarNotificacionError("error localidades");
      }
    },

    async guardar() {
      const cargaUtil = {
        curp: this.detalles.curp,
        noExpediente: this.detalles.noExpediente,
        nombre: this.detalles.NombreEmbarazada,
        FechaNacimiento: Utiles.obtenerCadenaFecha(
          this.detalles.fechaNacimiento
        ),
        domicilioReferencia: this.detalles.Direccion,
        localidad: this.localidadSeleccionada.text,
        municipio: this.municipioSeleccionado.text,
        telefono: this.detalles.TelefonoEmbarazada,
        lenguaIndigena: this.detalles.lenguaIndigena,
        emigro: this.detalles.emigro,
        derechohabiencia: this.detalles.Derechohabiencia,
        violencia: this.detalles.violencia,
        ConsultaPregestacional: Utiles.obtenerCadenaFecha(
          this.detalles.consultaPregestacional
        ),
        comorbilidades: this.detalles.comorbilidades,
        AsistenciaPreg: this.detalles.AsistenciaPreg,
        gestas: this.detalles.Gestas,
        paras: this.detalles.Paras,
        abortos: this.detalles.Abortos,
        cesareas: this.detalles.Cesareas,
        FechaUltimoEvento: Utiles.obtenerCadenaFecha(
          this.detalles.fechaUltimoEvento
        ),
        presentComplicaciones: this.detalles.presentComplicaciones,
        complicaciones: this.detalles.complicaciones,
        FechaUlmaMenstruacion: Utiles.obtenerCadenaFecha(
          this.detalles.fechaUlmaMenstruacion
        ),
        FechaProbableParto: Utiles.obtenerCadenaFecha(
          this.detalles.fechaProbableParto
        ),
        SGA: this.detalles.SGA,
        FechaConsulta: Utiles.obtenerCadenaFecha(this.detalles.fechaConsulta),
        noConsultasEmbarazo: this.detalles.noConsultasEmbarazo,
        noConsultasMes: this.detalles.noConsultasMes,
        rubeola: this.detalles.rubeola,
        FechaVacunaTDPrimera: Utiles.obtenerCadenaFecha(
          this.detalles.fechaVacunaTDPrimera
        ),
        FechaVacunaTDSegunda: Utiles.obtenerCadenaFecha(
          this.detalles.fechaVacunaTDSegunda
        ),
        FechaVacunaTDRefuerzo: Utiles.obtenerCadenaFecha(
          this.detalles.fechaVacunaTDRefuerzo
        ),
        FechaVacunaTDPA: Utiles.obtenerCadenaFecha(
          this.detalles.fechaVacunaTDPA
        ),
        FechaVacunaInfluenza: Utiles.obtenerCadenaFecha(
          this.detalles.fechaVacunaInfluenza
        ),
        vacunaCOVID: this.detalles.vacunaCOVID,
        grupoRH: this.detalles.grupoRH,
        ego: this.detalles.ego,
        BiometriaHematica: this.detalles.BiometriaHematica,
        leucocitos: this.detalles.leucocitos,
        plaquetas: this.detalles.plaquetas,
        vdrl: this.detalles.vdrl,
        vih: this.detalles.vih,
        estadoGlucosa: this.detalles.estadoGlucosa,
        resultadoGlucosa: this.detalles.resultadoGlucosa,
        caracteristicasFetls: this.detalles.caracteristicasFetls,
        malformaciones: this.detalles.malformaciones,
        LiquidoAmiotico: this.detalles.liquidoAmiotico,
        placenta: this.detalles.placenta,
        FechaProbableUSG: Utiles.obtenerCadenaFecha(
          this.detalles.fechaProbableUSG
        ),
        referencia: this.detalles.referencia,
        acudio_refe: this.detalles.acudio_refe,
        contrareferencia: this.detalles.contrareferencia,
        planSeguridad: this.detalles.planSeguridad,
        lugarParto: this.detalles.lugarParto,
        quienAtenderaParto: this.detalles.quienAtenderaParto,
        transporteAME: this.detalles.transporteAME,
        FechaEvento: this.detalles.fechaEvento,
        observaciones: this.detalles.observaciones,
        diagnostico: this.detalles.diagnostico,
      };

      const respuesta = await EmbarazadaService.insertEmbarazada(cargaUtil);
      if (respuesta !== true) {
        console.log("respuesta save", respuesta);
        DialogosService.mostrarError("registro con errores");
      } else {
        DialogosService.mostrarExitoso(
          "Datos guardados correctament en la base de datos"
        );
        this.detalles = {
          noExpediente: "",
          NombreEmbarazada: "",
          curp: "",
          Direccion: "",
          Derechohabiencia: "",
          TelefonoEmbarazada: null,
          DondeMigro: 0,
          fechaNacimiento: null,
          Gestas: 0,
          Paras: 0,
          Abortos: 0,
          Cesareas: 0,
          consultaPregestacional: null,
          fechaUltimoEvento: null,
          fechaUlmaMenstruacion: null,
          fechaProbableParto: null,
          fechaConsulta: null,
          fechaVacunaTDPrimera: null,
          fechaVacunaTDSegunda: null,
          fechaVacunaTDRefuerzo: null,
          fechaVacunaTDPA: null,
          fechaVacunaInfluenza: null,
          fechaProbableUSG: null,
          fechaEvento: null,
          edad: "",
          lenguaIndigena: "",
          emigro: 0,
          violencia: "",
          comorbilidades: "",
          AsistenciaPreg: 0,
          presentComplicaciones: 0,
          complicaciones: "",
          SGA: "",
          noConsultasEmbarazo: 0,
          noConsultasMes: 0,
          rubeola: 0,
          vacunaCOVID: 0,
          grupoRH: "",
          ego: "",
          BiometriaHematica: "",
          leucocitos: "",
          plaquetas: "",
          vdrl: "",
          vih: "",
          estadoGlucosa: "",
          resultadoGlucosa: "",
          caracteristicasFetls: "",
          malformaciones: "",
          liquidoAmiotico: "",
          placenta: "",
          referencia: "",
          acudio_refe: "",
          contrareferencia: "",
          planSeguridad: "",
          lugarParto: "",
          quienAtenderaParto: "",
          transporteAME: "",
          observaciones: "",
          diagnostico: "",
        };
      }
    },
  },
};
</script>

<style scoped>
.container {
  display: flex;
  flex-wrap: wrap;
  flex-direction: column;
  width: fit-content;
  height: fit-content;
  align-content: space-between;
}
</style>
