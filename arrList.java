 public class MyArrayList {

	int SIZE_FACTOR=5;
	
	Arr data[];
	
	int index;
	
	int size;
	
	public MyArrayList(){
		data=new Arr[SIZE_FACTOR];
		size=SIZE_FACTOR;
	}
	
	public void add(Arr obj){
		if(index==size-1){
        size=size+SIZE_FACTOR;
		Arr newData[]=new Arr[size];
		for(int i=0; i<data.length;i++){
			newData[i]=data[i];
		}
		data=newData;
        }
		data[index]=obj;
		index++;
	}
	
	public Arr get(int i){
		if(i>index-1){
			return 
		}
		if(i<0){
			return
		}
		return data[i];
		
	}
	
	public void remove(int i){
		if(i>index-1){
			return 
		}
		if(i<0){
			return
		}
		System.out.println("Arr getting removed:"+data[i]);
		for(int x=i; x<data.length-1;x++){
			data[x]=data[x+1];
		}
		index--;
	}

	public static void main(String[] args) {
		MyArrayList mal = new MyArrayList();
		mal.add("0");
		mal.add("1");
		mal.add("2");
		mal.add("3");
		mal.add("4");
		mal.add("5");
		mal.add("6");
		mal.add("7");
		mal.add("8");
		mal.add("9");
		
		mal.remove(5);
		System.out.println(mal.get(7));
	}
}
