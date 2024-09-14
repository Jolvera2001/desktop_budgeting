export namespace users {
	
	export class User {
	    _id: number;
	    email: string;
	    name: string;
	    budget_period?: number;
	
	    static createFrom(source: any = {}) {
	        return new User(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this._id = source["_id"];
	        this.email = source["email"];
	        this.name = source["name"];
	        this.budget_period = source["budget_period"];
	    }
	}
	export class UserDto {
	    email: string;
	    name: string;
	    budget_period: number;
	
	    static createFrom(source: any = {}) {
	        return new UserDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.email = source["email"];
	        this.name = source["name"];
	        this.budget_period = source["budget_period"];
	    }
	}

}

