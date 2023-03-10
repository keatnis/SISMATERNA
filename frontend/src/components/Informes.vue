<template>
    <!-- <div class="columns"> -->
    <b-steps v-model="activeStep" :animated="isAnimated" :rounded="isRounded" :has-navigation="hasNavigation"
        :icon-prev="prevIcon" :label-position="labelPosition" :mobile-mode="mobileMode">


        <div class="column has-text-centered">
            <article class="panel is-warning centered">
                <p class="panel-heading">
                    Informes
                </p>
            </article>
            <h3 class="is-size-3">Busca embarazadas por:</h3>
        </div>
        <div>
            <div class="columns">
                <div class="column">
                    <b-field label="Buscar por:">
                        <b-select placeholder="Buscar por:" expanded>
                            
                            <option value="Municipio">Municipio</option>
                            <option value="Region">Region</option>

                        </b-select>
                    </b-field>



                    <b-field class="radio">
                        <input type="radio" name="migrante" />
                        Migrantes
                    </b-field>
                    <b-field>
                    <label class="radio">
                        <input type="radio" name="mayor20" />
                        Menor de 20 años
                        <input type="radio" name="mayor20" />
                        Mayo de 35 años
                    </label>
                    </b-field>


                    <b-field label="Comorbilidades">
                        <b-select placeholder="Seleccionar" expanded>
                            <option value="Sobrepeso">Sobrepeso</option>
                            <option value="Obesida">Obesida</option>
                            <option value="Diabetes Mellitus">Diabetes Mellitus</option>
                            <option value="Hipertensión arterial">Hipertensión arterial</option>
                            <option value="Cardiopatía">Cardiopatía</option>
                            <option value="Epilepsia">Epilepsia</option>
                            <option value="Sifilis">Sifilis</option>
                            <option value="VIH">VIH</option>
                            <option value="Hepatitis">Hepatitis</option>
                            <option value="Artritis reumatoide">Artritis reumatoide</option>
                            <option value="Lupus">Lupus</option>
                            <option value="Cancer">Cancer</option>
                            <option value="Tuberculosis">Tuberculosis</option>
                            <option value="Covid-19">Covid-19</option>
                            <option value="Infección de vías urinarias">Infección de vías urinarias</option>
                            <option value="Cervicovaginitis">Cervicovaginitis</option>
                            <option value="Anemia">Anemia</option></b-select>
                    </b-field>

                    <b-field label="Número de gestas">
                        <b-select placeholder="Gestas" expanded>
                            <option value="Mephaa">5-7</option>
                            <option value="Amuzgo">8-9</option>
                        </b-select>
                    </b-field>
                    <b-field label="Trimmestre">
                        <label class="radio">
                            <input type="radio" name="trimestre" />
                            1Er Trimestre
                        </label>
                        <label class="radio">
                            <input type="radio" name="trimestre" />
                            2Do Trimestre
                        </label>
                        <label class="radio">
                            <input type="radio" name="trimestre" />
                            3Er Trimestre
                        </label>
                    </b-field>
                </div>



                <div class="column">
                    <b-field label="Semana de gestión">
                        <b-select placeholder="Semana de gestión" expanded>
                            <option value="Nahual">Nahual</option>
                            <option value="Mixteco">Tu´un Savi (Mixteco)</option>
                            <option value="Mephaa">Meephaa (Tlapaneco)</option>
                            <option value="Amuzgo">Amuzgo</option>
                            <option value="Ninguno">Ninguno</option>
                        </b-select>
                    </b-field>

                    <b-field label="VIH">
                        <label class="radio">
                            <input type="radio" name="react" />
                            Reactivo
                        </label>
                        <label class="radio">
                            <input type="radio" name="react" />
                            No reactivo
                        </label>
                    </b-field>

                    <b-field label="VDRL">
                        <label class="radio">
                            <input type="radio" name="vdrl" />
                            Positivo (+)
                        </label>
                        <label class="radio">
                            <input type="radio" name="vdrl" />
                            Negativo (-)
                        </label>
                    </b-field>

                    <b-field label="Grupo RH">
              <b-select placeholder="Grupo RH"  expanded>
                <option value="A+">A+</option>
                <option value="A-">A-</option>
                <option value="B+">B+</option>
                <option value="B-">B-</option>
                <option value="O+">O+</option>
                <option value="O-">O-</option>
                <option value="AB+">AB+</option>
                <option value="AB-">AB-</option>
                <option value="Ninguno">Ninguno</option>
              </b-select>
            </b-field>

                    <b-field label="Referencia">
                        <label class="radio">
                            <input type="radio" name="referen" />
                            Si
                        </label>
                        <label class="radio">
                            <input type="radio" name="referen" />
                            No
                        </label>
                    </b-field>

                    <b-field label="Proximas a parir">
                        <b-select placeholder="Seleccionar" expanded>
                            <option value="15 Dias">15 Dias</option>
                            <option value="1 Mes">1 Mes</option>
                        </b-select>
                    </b-field>
                    <b-table :data="embarazadas" :columns="columns" :sticky-header="stickyHeaders" ></b-table>  
                    <div class="buttons">
                        <button class="button is-info">BUSCAR</button>
                    </div>
                </div>
            </div>
        </div>
    </b-steps>

</template>
<script>

</script>