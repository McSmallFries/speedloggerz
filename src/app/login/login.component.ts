import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from '../service/auth.service';
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  loginForm: FormGroup;
  message: string;
  returnUrl: string;


  constructor(private formBuilder: FormBuilder, private router: Router, public authService: AuthService) { }

  ngOnInit() {
    this.loginForm = this.formBuilder.group({
      username: ['', Validators.required],
      password: ['', Validators.required]
    });
  }

  // convenience getter for easy access to form fields
  get getLoginFormControls() { return this.loginForm.controls; }

  signUp()  {
    this.router.navigate(['/register'])
  }

  login() {
    debugger;
    if (this.loginForm.invalid) {
        return;
    }
    if (this.getLoginFormControls.username.value !== '' && this.getLoginFormControls.password.value !== '') {
      // console.log('Login successful');
      // this.authService.authLogin(this.model);

      localStorage.setItem('isLoggedIn', 'true');
      localStorage.setItem('token', this.getLoginFormControls.username.value);
      this.router.navigate(['/dashboard']);
    } else {
      this.message = 'Please check your username and password';
    }
  }
}
