import { CommonModule } from '@angular/common';
import { Component, Input } from '@angular/core';
import {
  IonBadge,
  IonButton,
  IonButtons,
  IonIcon,
  IonItem,
  IonItemDivider,
  IonItemGroup,
  IonLabel,
} from '@ionic/angular/standalone';
import { addIcons } from 'ionicons';
import { open } from 'ionicons/icons';
import { CarbonModule } from 'kcarbon';
import escape from 'escape-html';

import { pb } from 'src/pb';

@Component({
  selector: 'app-ack',
  templateUrl: './ack.component.html',
  styleUrls: ['./ack.component.scss'],
  standalone: true,
  imports: [
    IonItem,
    IonItemGroup,
    IonLabel,
    IonItemDivider,
    IonBadge,
    CommonModule,
    CarbonModule,
    IonButton,
    IonButtons,
    IonIcon,
  ],
})
export class AckComponent {
  @Input()
  item?: pb.IAck;
  constructor() {
    addIcons({ open });
  }

  html(mate: pb.IMate) {
    if (!mate.hits || !mate.text) {
      return '';
    }

    let start = 0;

    const list = [];
    for (const hit of mate.hits) {
      if (!hit.len) {
        continue;
      }

      if (!hit.col) {
        hit.col = 0;
      }

      list.push(escape(mate.text.substring(start, hit.col)));
      list.push('<b>');
      list.push(escape(mate!.text.substring(hit.col, hit.col + hit.len)));
      list.push('</b>');
      start = hit.col + hit.len;
    }

    list.push(escape(mate.text.substring(start)));

    return list.join('');
  }
}
