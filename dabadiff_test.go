package main

import (
	"testing"  
)

func TestDiff(t *testing.T) {
	 
	diffs := Diffs()
		
	if len(diffs) > 0 {
		for tt,ds := range diffs {			
			t.Errorf("%s has diff:\n%s",tt,ds)
		}				
	} else {
		t.Log("Pass")
	}
	
} 