import { Component, OnInit } from '@angular/core';
import { AppGlobalService } from './service/globals.service';
import { SpeedloggerzService } from './service/speedloggerz.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit{
  title = 'SpeedLoggerz';

  constructor(private globals: AppGlobalService, private service: SpeedloggerzService, private router: Router)  {
    
  }

  ngOnInit()  {
    this.globals.setIsUserLoggedIn(localStorage.getItem('isLoggedIn') === 'true')
    if (this.globals.isUserLoggedIn)  {
      this.router.navigate(['/dashboard']);
      return;
    } 
    else {
      this.router.navigate(['/login']);
      return;
    }
  }
}
