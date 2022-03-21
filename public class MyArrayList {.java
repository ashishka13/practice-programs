public class MyArrayList {

	private static final int initialSize=5; // give the size for first time
	
	private int data[]; // data array of integer type
	
	private int index;
	
	private int size;
	
	public MyArrayList(){   // initialize object with initial array size
		data=new int[initialSize];  // 
		size=initialSize;
	}
	
	public void add(int num){
		if(index == size-1){    // check if the array is getting full, if yes then increase size
			size=size+initialSize;
		
            int newData[]=new int[size];    // increase size
		
            for(int i=0; i<data.length;i++){
			newData[i]=data[i];
		}
		data=newData;
		}
		data[index]=num;
		index++;    // increment index 
	}
	
	
	public static void main(String[] args) throws Exception {
		MyArrayList myArrl = new MyArrayList(); // create object to call methods on
		myArrl.add("0");    // call add method
		myArrl.add("1");
		myArrl.add("2");
		myArrl.add("3");	
	}

}
