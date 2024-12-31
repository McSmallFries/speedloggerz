import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from '../service/auth.service';
import {SpeedloggerzService} from '../service/speedloggerz.service';
import { User } from '../models/login';
@Component({
  selector: 'app-login',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css']
})
export class SignupComponent implements OnInit {
  signupForm: FormGroup;
  message: string;
  returnUrl: string;


  constructor(private formBuilder: FormBuilder, private router: Router, public authService: AuthService, private appservice: SpeedloggerzService) { }

  ngOnInit() {
    this.signupForm = this.formBuilder.group({
      username: ['', Validators.required],
      password: ['', Validators.required]
    });
    if (localStorage.getItem('isLoggedIn') === 'true') {
      this.router.navigate(['/dashboard']);
      return;
    } else {
      this.router.navigate(['/register']);
      return;
    }
    // this.returnUrl = '/dashboard';
    // this.authService.logout();
  }

  // convenience getter for easy access to form fields
  get getFormControls() { return this.signupForm.controls; }

  async register() {
    if (this.signupForm.invalid) {
        return;
    } else if (this.getFormControls.username.value !== '' && this.getFormControls.password.value !== '') {
      // console.log('Login successful');
      // this.authService.authLogin(this.model);
      // this.
      // localStorage.setItem('isLoggedIn', 'true');
      // localStorage.setItem('token', this.getFormControls.username.value);
      // this.router.navigate([this.returnUrl]);
      const usr = new User();
      usr.Username = this.getFormControls.username.value;
      usr.Password = this.getFormControls.password.value;
      await this.appservice.RegisterUser(usr);

    } else {
      this.message = 'Please check your username and password';
    }
  }
}
