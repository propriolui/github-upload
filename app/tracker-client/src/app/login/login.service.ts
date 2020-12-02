import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/internal/Observable';
import { HttpHeaders } from '@angular/common/http';
import { Subscription } from 'rxjs';

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type': 'application/json'
    })
};

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  readonly ROOT_URL = 'http://localhost:8080/login';

 

  response: Subscription | undefined;

  constructor(private http: HttpClient) { }
  

  public login(email: string, pwd: string){
    var credentials = JSON.stringify({email: email, password: pwd});
    const headers = { 'content-type': 'application/json'}  
    console.log(credentials);
    this.response = this.http.post('http://localhost:8080/login', credentials, httpOptions).subscribe(x => console.log(x));
    console.log(this.response)
  }
}
