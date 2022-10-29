export namespace internal {
	
	export class MdPath {
	    id: number;
	    path: string;
	    type: number;
	    fname: string;
	    // Go type: time.Time
	    created: any;
	
	    static createFrom(source: any = {}) {
	        return new MdPath(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.path = source["path"];
	        this.type = source["type"];
	        this.fname = source["fname"];
	        this.created = this.convertValues(source["created"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class TodoItem {
	    id: number;
	    title: string;
	    content: string;
	    tags: string;
	    // Go type: time.Time
	    date: any;
	    hasContent: boolean;
	    done: boolean;
	    importent: number;
	    // Go type: time.Time
	    expired: any;
	    // Go type: time.Time
	    updated: any;
	
	    static createFrom(source: any = {}) {
	        return new TodoItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.content = source["content"];
	        this.tags = source["tags"];
	        this.date = this.convertValues(source["date"], null);
	        this.hasContent = source["hasContent"];
	        this.done = source["done"];
	        this.importent = source["importent"];
	        this.expired = this.convertValues(source["expired"], null);
	        this.updated = this.convertValues(source["updated"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace util {
	
	export class Resp {
	    code: number;
	    msg: string;
	    data: any;
	
	    static createFrom(source: any = {}) {
	        return new Resp(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.msg = source["msg"];
	        this.data = source["data"];
	    }
	}

}

