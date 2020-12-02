import { Component, NgModule, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { LoginService } from './login.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  minLength = 8;
  hide = true;
  email = new FormControl('', [Validators.required, Validators.email]);
  pwd = new FormControl('', [Validators.required, Validators.minLength(this.minLength)]);

  form = new FormGroup({
    email: this.email,
    pwd: this.pwd
  })
 

  constructor(private login: LoginService, private router: Router) {}

  ngOnInit(): void{
  }            

  getEmailError(){
    if (this.form.controls["email"].hasError("required")){
      return "Campo richiesto";
    }

    if (this.form.controls["email"].hasError("email")){
      return "Formato non valido";
    }
    return
  }

  getPwdError(){
    if (this.form.controls["pwd"].hasError("required")){
      return "Campo richiesto";
    }
    if (this.form.controls["pwd"].hasError("minlength")){
      return "Lunghezza minima "+ this.minLength + " caratteri";
    }
    return
  }


  onSubmit(){
    const val = this.form.value    
    if(!this.email.valid || !this.pwd.valid){
      return
    }
    if(val.email && val.pwd) {

      this.login.login(val.email, val.pwd);
      
    }
  }

}
