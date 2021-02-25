/*
 * Copyright © 2017. TIBCO Software Inc.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */
import { NgModule } from "@angular/core";
import { CommonModule } from "@angular/common";
import { TibcoDatagridConnectorContribution } from "./connector";
import { WiServiceContribution } from "wi-studio/app/contrib/wi-contrib";

@NgModule({
  imports: [CommonModule],
  providers: [
    {
      provide: WiServiceContribution,
      useClass: TibcoDatagridConnectorContribution,
    },
  ],
})
export default class TibcoDatagridConnectorModule {}
