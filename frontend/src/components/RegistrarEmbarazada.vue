<template>
  <!-- <div class="columns"> -->
  <b-steps
    v-model="activeStep"
    :animated="isAnimated"
    :rounded="isRounded"
    :has-navigation="hasNavigation"
    :icon-prev="prevIcon"
    :icon-next="nextIcon"
    :mobile-mode="mobileMode"
  >
    <b-step-item step="" label="" :clickable="isStepsClickable">
      <div class="column has-text-centered">
        <article class="panel is-info centered">
          <p class="panel-heading">Registrar embarazada</p>
        </article>
      </div>
      <div>
        <div class="columns">
          <div class="column is-one-third">
            <b-field label="No. Progreso">
              <b-input disabled v-model="detalles.id"></b-input>
            </b-field>

            <b-field label="Nombre">
              <b-input
                v-model="detalles.NombreEmbarazada"
                placeholder="Nombre completo de paciente"
              ></b-input>
            </b-field>

            <b-field label="Número  teléfono">
              <b-input
                v-model="detalles.TelefonoEmbarazada"
                placeholder="Teléfono fijo"
              ></b-input>
            </b-field>
            <b-field label="Municipio">
              <b-select
                placeholder="Seleccionar municipio"
                v-model="municipioSeleccionado"
                :onChange="obtenerLocalidades()"
                expanded
              >
                <option
                  v-for="option in municipios"
                  :value="{
                    id: option.id_municipio,
                    text: option.nombreMunicipio,
                  }"
                  :key="option.id"
                >
                  {{ option.nombreMunicipio }}
                </option>
              </b-select>
            </b-field>

            <b-field label="Derechohabiencia">
              <b-select v-model="detalles.Derechohabiencia" expanded>
                <option value="">Seleccionar</option>
                <option value="Insabi">INSABIi</option>
                <option value="Imss">IMSS</option>
                <option value="Issste">ISSSTE</option>
                <option value="Sedena">SEDENA</option>
                <option value="Pemex">PEMEX</option>
                <option value="Marina">MARINA</option>
                <option value="Otro">Otro</option>
              </b-select>
            </b-field>
            <b-field label="Detenciones">
              <b-input
                v-model="detalles.detenciones"
                placeholder="Detenciones"
              ></b-input>
            </b-field>
            <b-field label="Comorbilidades">
              <section>
                <div class="container">
                  <b-checkbox
                    v-model="detalles.sobrepeso"
                    native-value="sobrepeso"
                  >
                    Sobre peso
                  </b-checkbox>
                  <b-checkbox
                    v-model="detalles.obesidad"
                    native-value="Obesidad"
                  >
                    Obesidad
                  </b-checkbox>
                  <b-checkbox
                    v-model="detalles.diabetesMellitus"
                    native-value="Diabetes Mellitus"
                  >
                    Diabetes Mellitus
                  </b-checkbox>
                  <b-checkbox
                    v-model="detalles.hipertensionArterial"
                    native-value="Hipertensión arterial"
                  >
                    Hipertensión arterial
                  </b-checkbox>
                  <b-checkbox
                    v-model="detalles.cardeopatia"
                    native-value="Cardiopatía"
                  >
                    Cardiopatía
                  </b-checkbox>
                  <b-checkbox
                    v-model="detalles.Epilepsia"
                    native-value="Epilepsia "
                  >
                    Epilepsia
                  </b-checkbox>
                  <b-checkbox
                    v-model="detalles.Sifilis"
                    native-value="Sifilis "
                  >
                    Sifilis
                  </b-checkbox>
                  <b-checkbox v-model="detalles.vih" native-value="VIH ">
                    VIH
                  </b-checkbox>
                  <b-checkbox
                    v-model="detalles.hepatitis"
                    native-value="Hepatitis"
                  >
                    Hepatitis
                  </b-checkbox>
                  <b-checkbox v-model="detalles.ninguno" native-value="Ninguno">
                    Ninguno
                  </b-checkbox>
                </div>
                <p class="content">
                  <b>Selección:</b>
                  {{ checkboxGroup }}
                </p>
              </section>
            </b-field>
            <b-field label="Gestas">
              
              <b-numberinput
                v-model="detalles.Gestas"
                type="number"

                placeholder="Número de gestas"
              ></b-numberinput>
            </b-field>
            <b-field label="Cesareas">
              <b-numberinput
                v-model="detalles.Cesareas"
                placeholder="Número de cesareas"
              ></b-numberinput>
            </b-field>
          </div>

          <div class="column is-one-third">
            <b-field label="CURP">
              <b-input
                v-model="detalles.curp"
                placeholder="Clave Unica de Registro de Población"
              ></b-input>
            </b-field>
            <b-field label="Fecha de nacimiento">
              <b-datepicker
                v-model="detalles.fechaNacimiento"
                placeholder="Fecha de nacimiento"
                icon="calendar-today"
                :datetime-formatter="formatearFecha"
                :locale="locale"
                @input="calcularEdad()"
                editable
              >
              </b-datepicker>
            </b-field>

            <b-field label="Lengua Indígena">
              <b-select
                v-model="detalles.lenguaIndigena"
                placeholder="Lengua Indígena"
                expanded
              >
                <option value="Nahual">Nahual</option>
                <option value="Mixteco">Tu´un Savi (Mixteco)</option>
                <option value="Mephaa">Meephaa (Tlapaneco)</option>
                <option value="Amuzgo">Amuzgo</option>
                <option value="Ninguno">Ninguno</option>
              </b-select>
            </b-field>
            <b-field label="Localidad">
              <b-select
                placeholder="Seleccionar localidad"
                v-model="localidadSeleccionada"
                expanded
              >
                <option
                  v-for="option in localidades"
                  :value="{
                    id: option.id_localidad,
                    text: option.NombreLocalidad,
                  }"
                  :key="option.id"
                >
                  {{ option.NombreLocalidad }}
                </option>
              </b-select>
            </b-field>

            <b-field label="De tenciones"></b-field>
            <div class="field">
              <div class="b-checkbox">
                <input
                  type="checkbox"
                  name="Mamanormal"
                  value="verde"
                  id="MamaNormal"
                />

                <label for="MamaNormal">
                  Exploración clínica de mama normal
                </label>
              </div>
              <div class="b-checkbox">
                <input
                  id="MamaAlterado"
                  class="styled"
                  checked
                  type="checkbox"
                />
                <label for="MamaAlterado">
                  Exploración clínica de mama alterado
                </label>
                <div class="b-checkbox">
                  <input
                    id="PapaNormal"
                    class="styled"
                    checked
                    type="checkbox"
                  />
                  <label for="PapaNormal"> Papanicolau normal </label>
                </div>
                <div class="b-checkbox">
                  <input
                    id="PapaLesion"
                    class="styled"
                    checked
                    type="checkbox"
                  />
                  <label for="PapaLesion"> Papanicolau con lesiones </label>
                </div>
                <div class="b-checkbox">
                  <input
                    id="PapaCancer"
                    class="styled"
                    checked
                    type="checkbox"
                  />
                  <label for="PapaCancer"> Papanicolau cancer IN SITU </label>
                </div>

                <div class="b-checkbox">
                  <input
                    id="PapaInvasor"
                    class="styled"
                    checked
                    type="checkbox"
                  />
                  <label for="PapaInvasor"> Papanicolau invasor </label>
                </div>
                <div class="b-checkbox">
                  <input id="NoDete" class="styled" checked type="checkbox" />
                  <label for="NoDete"> No se realizo </label>
                </div>
              </div>

              <b-field label="Paras">
                <b-numberinput
                  v-model="detalles.Paras"
                  type="number"
                  placeholder="Numero de paras"
                ></b-numberinput>
              </b-field>
              <b-field label="Abortos">
                <b-numberinput
                  v-model="detalles.Abortos"
                  type="number"
                  placeholder="Numero de abortos"
                ></b-numberinput>
              </b-field>
            </div>
          </div>
          <div class="column is-one-fifthr">
            <b-field label="No. Expediente clínico">
              <b-input
                v-model="detalles.noExpediente"
                placeholder="No. Expediente clínico"
              ></b-input>
            </b-field>
            <b-field label="Edad">
              <b-input
                disabled="»disabled»"
                v-model="detalles.edad"
                value="this.deatalles.edad"
              ></b-input>
            </b-field>
            <b-field label="Emigro">
              <b-radio v-model="detalles.emigro" :native-value="false"
                >Si</b-radio
              >
              <b-radio v-model="detalles.emigro" :native-value="true"
                >No</b-radio
              >
            </b-field>

            <b-field label="Dirección">
              <b-input
                v-model="detalles.Direccion"
                placeholder="Dirección con referencia"
              ></b-input>
            </b-field>
            <div class="control">
              <b-field label="¿Asistió a consulta pregestacional?"></b-field>
              <label class="radio">
                <input type="radio" name="ConsultaPreg" />
                Si
              </label>
              <label class="radio">
                <input type="radio" name="ConsultaPreg" />
                No
              </label>
            </div>

            <b-datepicker
              v-model="detalles.consultaPregestacional"
              placeholder="Fecha de consulta"
              icon="calendar-today"
              :datetime-formatter="formatearFecha"
              :locale="locale"
              editable
            >
            </b-datepicker>

            <b-field label="Fecha de ultimo evento( Parto, cesarea, aborto)">
              <b-datepicker
                v-model="detalles.fechaUltimoEvento"
                placeholder="Fecha de ultimo evento"
                icon="calendar-today"
                :datetime-formatter="formatearFecha"
                :locale="locale"
                editable
              >
              </b-datepicker>
            </b-field>

            <div class="control">
              <b-field
                label="¿Presentó complicaciones en el embarazo previo?"
              ></b-field>
              <label class="radio">
                <input type="radio" name="Complicaciones" />
                Si
              </label>
              <label class="radio">
                <input type="radio" name="Complicaciones" />
                No
              </label>
              <b-select
                placeholder="¿Qué complicaciones presento?"
                v-model="detalles.complicaciones"
                expanded
              >
                <option value="Hemorragia ">Hemorragia</option>
                <option value="RetencioPlacenta">Retención de placenta</option>
                <option value="PlacentaPrevia">Placenta previa</option>
                <option value="DiabetesGesta">Diabetes gestacional</option>
                <option value="Preeclampsia">Preeclampsia</option>
                <option value="Eclampsia">Eclampsia</option>
                <option value="CovidEmbarazo">Covid + embarazo</option>
                <option value="OtraCompli">Otra</option>
                <option value="NingunaCompli">Ninguna</option>
              </b-select>
            </div>
          </div>
        </div>
      </div>
    </b-step-item>

    <b-step-item
      step=""
      label=""
      :clickable="isStepsClickable"
      :type="{ 'is-success': isProfileSuccess }"
    >
      <div>
        <div class="columns">
          <div class="column is-one-third">
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

            <div class="control">
              <b-field label="Fecha de consutal">
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

              <b-field label="Rubeola">
                <label class="radio">
                  <input type="radio" name="Rubeola" />
                  Si se aplicó previo al embarazo
                </label>
                <label class="radio">
                  <input type="radio" name="Rubeola" />
                  No se cuenta con vacuna
                </label>
              </b-field>

              <b-field label="Fecha de vacunación TD refuerzo">
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

              <b-field label="Vacuna COVID">
                <b-radio v-model="detalles.vacunaCOVID" :native-value="1"
                  >Se aplicó</b-radio
                >
                <b-radio v-model="detalles.vacunaCOVID" :native-value="0"
                  >No se aplico</b-radio
                >
              </b-field>

              <b-field label="Biometria ematica">
                <b-select placeholder="Biometria" expanded>
                  <option value="HB12">HB >=12</option>
                  <option value="HBMAY9">HB >=9</option>
                  <option value="HBMEN9">HB Menor a 9</option>
                  <option value="Ninguno">Ninguno</option>
                </b-select>
              </b-field>

              <b-field label="VIH">
                <b-select placeholder="VIH" expanded>
                  <option value="ReactivoVIH">Reactivo</option>
                  <option value="NoReactivoVIH">No reactivo</option>
                  <option value="NoRealizoVIH">No se realizó</option>
                </b-select>
              </b-field>
              <b-field label="Características fetales ">
                <b-select placeholder="Características" expanded>
                  <option value="UnicoVivo">Unico vivo</option>
                  <option value="UnicObito">Unico obito</option>
                  <option value="Molar">Molar</option>
                  <option value="Monocorial">Monocorial</option>
                  <option value="Bicorionicos">Bicorionicos</option>
                  <option value="Gemelar1Obitado">Gmelar 1 obitado</option>
                </b-select>
              </b-field>
            </div>
          </div>

          <div class="column is-one-third">
            <div class="control">
              <b-field label="Fecha probable de parto">
                <b-datepicker
                  v-model="detalles.fechaProbableParto"
                  placeholder="Fecha probable de parto"
                  icon="calendar-today"
                  :locale="locale"
                  editable
                >
                </b-datepicker>
              </b-field>
            </div>
            <b-field label="No. Total de consulta durante el embarazo">
              <b-input
                disabled="»disabled»"
                v-model="detalles.placas"
                placeholder="No. Total de consulta durante el embarazo"
              ></b-input>
            </b-field>
            <div class="control">
              <b-field label="Fecha de vacunación TD primera">
                <b-datepicker
                  v-model="detalles.fechaVacunaTDPrimera"
                  placeholder="Fecha de vacunación TD primera "
                  icon="calendar-today"
                  :locale="locale"
                  editable
                >
                </b-datepicker>
              </b-field>
            </div>
            <div class="control">
              <b-field label="Fecha de vacunación TDPA">
                <b-datepicker
                  v-model="detalles.fechaVacunaTDPA"
                  placeholder="Fecha de vacunación TDPA "
                  icon="calendar-today"
                  :locale="locale"
                  editable
                >
                </b-datepicker>
              </b-field>
            </div>

            <b-field label="Grupo RH">
              <b-select placeholder="Grupo RH" expanded>
                <option value="A+">A+</option>
                <option value="A-">A-</option>
                <option value="B+">B+</option>
                <option value="B-">B-</option>
                <option value="O+">O+</option>
                <option value="O-">O-</option>
                <option value="AB+">AB+</option>
                <option value="AB-">AB-</option>
                <option value="NingunoRH">Ninguno</option>
              </b-select>
            </b-field>
            <b-field label="Leucocitos">
              <b-select placeholder="Leucocitos" expanded>
                <option value="Normales">Normales</option>
                <option value="DisminuidosLinfocitos">
                  Disminuidos linfocitos
                </option>
                <option value="AumentadosNeutrofilos">
                  Aumentados neutrofilos
                </option>
                <option value="AumentadosLinfocitos">
                  Aumentados los linfocitos
                </option>
                <option value="NingunoLEUCO">Ninguno</option>
              </b-select>
            </b-field>

            <b-field label="Estado tiras glucosa">
              <b-select placeholder="Estado en el que se realizó" expanded>
                <option value="Ayuno">Ayuno</option>
                <option value="Posprandial1">Posprandial 1 hora</option>
                <option value="Posprandial2">Posprandial 2 horas</option>
                <option value="NingunoLEUCO">Ninguno</option>
              </b-select>
            </b-field>

            <b-field label=" Malformaciones congenitas">
              <b-select placeholder="Malformaciones" expanded>
                <option value="EspinaBifida">Espina bifida</option>
                <option value="Meningocele">Meningocele</option>
                <option value="Mielomeningocele">Mielomeningocele</option>
                <option value="Anencefalia">Anencefalia</option>
                <option value="OtrasMalform">Otras</option>
                <option value="NingunaMalform">Ninguna</option>
              </b-select>
            </b-field>
          </div>
          <div class="column is-one-third">
            <b-field label="Semana de gestación actual">
              <b-input
                disabled="»disabled»"
                v-model="detalles.semanaGestacion"
                placeholder="Semanan de gestación actual"
              >
              </b-input>
            </b-field>
            <b-field label="No. Total de consulta en el mes">
              <b-input
                disabled="»disabled»"
                v-model="detalles.placas"
                placeholder="No. Total de consulta en el mes "
              >
              </b-input>
            </b-field>
            <div class="control">
              <b-field label="Fecha de vacunación TD segunda">
                <b-datepicker
                  v-model="detalles.fechaVacunaTDSegunda"
                  placeholder="Fecha de vacunación TD segunda "
                  icon="calendar-today"
                  :datetime-formatter="formatearFecha"
                  :locale="locale"
                  editable
                >
                </b-datepicker>
              </b-field>
            </div>
            <div class="control">
              <b-field label="Fecha de vacunación Influenza estacional">
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
            </div>
            <b-field label="EGO y/o tiras urianalisis">
              <b-select placeholder="EGO" expanded>
                <option value="Trazas">Prot. trazas</option>
                <option value="Mayor300">Prot. >300 mg/l</option>
                <option value="Leucitos">leucitos >5</option>
                <option value="Nitratos">nitratos +</option>
                <option value="NegativoEgo">Negativo</option>
              </b-select>
            </b-field>

            <b-field label="VDRL">
              <b-select placeholder="VDRL" expanded>
                <option value="ReactivoVDRL">Reactivo</option>
                <option value="NoReactivoVDRL">No reactivo</option>
                <option value="NoRealizo">No se realizo</option>
              </b-select>
            </b-field>
            <b-field label="Resultado tiras glucosa">
              <b-select placeholder="Resultado tiras glucosa" expanded>
                <option value="Mayor60">>60 y menor a 110</option>
                <option value="Mayor110">>=110 y menor o igual a 126</option>
                <option value="Mayor126">>126 y menor a 160</option>
                <option value="NingunoGlucosa">Ninguno</option>
              </b-select>
            </b-field>

            <b-field label="Liquido amniotico">
              <b-select placeholder="Liquido amniotico" expanded>
                <option value="Normal">Normal</option>
                <option value="Oligohidramnios">Oligohidramnios</option>
                <option value="Polihidramnios">Polihidramnios</option>
                <option value="Anhidramnios">Anhidramnios</option>
              </b-select>
            </b-field>
          </div>
        </div>
      </div>
    </b-step-item>

    <b-step-item
      step=""
      label=""
      :clickable="isStepsClickable"
      :type="{ 'is-success': isProfileSuccess }"
    >
      <div>
        <div class="columns">
          <div class="column is-one-third">
            <b-field label="Placenta">
              <b-select placeholder="Placenta" expanded>
                <option value="NormalPlacen">Normal</option>
                <option value="Previa">Previa</option>
                <option value="AcretismoPlacentario">
                  Acretismo placentario
                </option>
                <option value="Calcificada">Calcificada</option>
              </b-select>
            </b-field>

            <b-field label="¿Acudió a referencia?"></b-field>
            <b-select placeholder="Referencias" expanded>
              <option value="HgTlapa ">H.G. Tlapa</option>
              <option value="Hmnig">HMNIG</option>
              <option value="HcAcatepec">H.C. Acatepec</option>
              <option value="HcAlcozauca">H.C. Alcozauca</option>
              <option value="HcMalina">H.C. Malinaltepec</option>
              <option value="HcOlinala">H.C. Olinalá</option>
              <option value="HcTlacoapa">H.C. Tlacoapa</option>
              <option value="HcXochi">H.C. Xochihuehuetlán</option>
              <option value="HcZapo">H.C. Zapotitlán tablas</option>
              <option value="CentroSalud">Centro de salud</option>
              <option value="Hogar">HOGAR</option>
              <option value="HcSaLuis">H.C. San sluis</option>
              <option value="HgAyutla">H.G. Ayutla</option>
              <option value="HgOme">H.G. Ometepec</option>
              <option value="HgChuautla">H.G Chuautla</option>
              <option value="OtrosHospitales">Otros hospitales</option>
              <option value="NoAcudio">No acudió</option>
            </b-select>

            <b-field
              label="¿Se oriento en señales de peligro durante el embarazo, parto, puerperio, rn?"
            ></b-field>
            <b-select
              placeholder="¿Se oriento en señales de peligro durante el embarazo?"
              expanded
            >
              <option value="SIOriento">Si</option>
              <option value="NoOriento">No</option>
            </b-select>
            <b-field label="¿cuenta con transporte ame en caso de emergencia?">
              <b-select placeholder="¿Cuenta con transporte AME?" expanded>
                <option value="SiAME">SI</option>
                <option value="NoAME">NO</option>
              </b-select>
            </b-field>
          </div>

          <div class="column is-one-third">
            <div class="control">
              <b-field label="Fecha probable de parto por USG">
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
            </div>
            <b-field label="Contrareferencia recibida">
              <b-select placeholder="Referencia" expanded>
                <option value="SiConrefe">Si</option>
                <option value="NoCorefe">No</option>
              </b-select>
            </b-field>
            <b-field label="¿Donde atenderán su parto?">
              <b-select placeholder="¿Donde atenderán su parto?" expanded>
                <option value="AtenderHospital">Hospital</option>
                <option value="AtenderC.Salud">C. Salud</option>
                <option value="AtenderHogar">Hogar</option>
                <option value="AtenderParticular">Particular</option>
              </b-select>
            </b-field>
            <div class="control">
              <b-field label="Fecha de evento">
                <b-datepicker
                  v-model="detalles.fechaEvento"
                  placeholder="Fecha de evento"
                  icon="calendar-today"
                  :datetime-formatter="formatearFecha"
                  :locale="locale"
                  editable
                >
                </b-datepicker>
              </b-field>
            </div>
          </div>
          <div class="column is-one-third">
            <b-field label="Referencia">
              <b-select placeholder="Referencia" expanded>
                <option value="SiRefe">Si</option>
                <option value="NoRefe">No</option>
              </b-select>
            </b-field>

            <b-field label="Se le realizó el plan de seguridad">
              <b-select placeholder="Plan de seguridad" expanded>
                <option value="SiRefe">Si</option>
                <option value="NoRefe">Reforzamiento</option>
                <option value="NoRefe">No</option>
              </b-select>
            </b-field>

            <b-field label="¿Quien atenderá su parto?">
              <b-select placeholder="¿Quien atenderá su parto?" expanded>
                <option value="AtenderMedico">Medico</option>
                <option value="AtenderEnfra">Enfra</option>
                <option value="AtenderParteraProfs">Partera profes.</option>
                <option value="AtenderParteraTradc">Partera tradic.</option>
                <option value="AtenderFamiliar">Familiar</option>
                <option value="AtenderOtro">Otro</option>
              </b-select>
            </b-field>
          </div>
        </div>
      </div>
    </b-step-item>
    <b-step-item
      :step="showSocial ? '' : ''"
      label=""
      :clickable="isStepsClickable"
      disabled
    >
      <Triagge />
      <b-field>
        <b-button
          label="Terminar registro"
          type="is-info"
          size="is-medium"
          @click="confirm()"
        />
      </b-field>
    </b-step-item>
  </b-steps>
</template>
<script>
import Utiles from "../services/Utiles";
import DialogosService from "../services/DialogosService";
import EmbarazadaService from "../services/EmbarazadaService";
import { DialogProgrammatic as Dialog } from "buefy";
import triagge from "./Triagge.vue";
export default {
  components: {
    Triagge: triagge,
  },
  data: () => ({
    selected: null,
    edad: null,
    activeStep: 0,
    municipios: [],
    localidades: [],
    localidadSeleccionada: "",
    municipioSeleccionado: "",
    formatearFecha: Utiles.formatearFechaSegunLocale,
    fechaInicio: null,
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
      emigro: false,
      detenciones: "",
      comorbilidades: "",
      AsistenciaPreg: 0,
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
      signosAlarma: "",
      lugarParto: "",
      quienAtenderaParto: "",
      transporteAME: "",
      observaciones: "",
      diagnostico: "",
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
      //()  var dia = cumpleanos.getDay();

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
        domicilioReferencia: this.detalles.domicilioReferencia,
        localidad: this.detalles.localidadSeleccionada,
        municipio: this.detalles.municipioSeleccionado,
        telefono: this.detalles.TelefonoEmbarazada,
        lenguaIndigena: this.detalles.lenguaIndigena,
        emigro: 1,
        derechohabiencia: this.detalles.Derechohabiencia,
        detenciones: this.detalles.detenciones,
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
          this.detalles.FechaUltimoEvento
        ),
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
        signosAlarma: this.detalles.signosAlarma,
        lugarParto: this.detalles.lugarParto,
        quienAtenderaParto: this.detalles.quienAtenderaParto,
        transporteAME: this.detalles.transporteAME,
        FechaEvento: this.detalles.fechaEvento,
        observaciones: this.detalles.observaciones,
        diagnostico: this.detalles.diagnostico,
      };

      const respuesta = await EmbarazadaService.insertEmbarazada(cargaUtil);
      if (respuesta != true) {
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
          emigro: false,
          detenciones: "",
          comorbilidades: "",
          AsistenciaPreg: 0,
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
          signosAlarma: "",
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
  width: 500px;
  height: 200px;
  align-content: space-between;
}
</style>