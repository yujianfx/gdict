package gdict

import (
	"fmt"
	"testing"
)

func TestDict_dictAdd(t *testing.T) {
	d := NewDict()
	for i := 0; i < 4096; i++ {
		d.dictAdd(fmt.Sprintf("key%d", i), i)
	}
	d.dictPrintStats()
}

func TestDict_dictFind(t *testing.T) {
	d := NewDict()
	t.Run("dictAdd", func(t *testing.T) {
		for i := 0; i < 409600; i++ {
			d.dictAdd(fmt.Sprintf("key%d", i), i)
		}
	})
	d.dictPrintStats()
	t.Run("dictFind", func(t *testing.T) {
		for i := 0; i < 409600; i++ {
			find := d.dictFind(fmt.Sprintf("key%d", i))
			if find == nil {
				t.Errorf("dictFind failed, key: %s", fmt.Sprintf("key%d", i))
			}
			if find.value != i {
				t.Errorf("dictFind failed, key: %s, value: %d", fmt.Sprintf("key%d", i), find.value)
			}
		}
	})
	d.dictPrintStats()
}

func TestDict_dictDelete(t *testing.T) {
	d := NewDict()
	for i := 0; i < 40960; i++ {
		d.dictAdd(fmt.Sprintf("key%d", i), i)
	}
	d.dictPrintStats()
	t.Run("dictDelete", func(t *testing.T) {
		for i := 0; i < 40960; i++ {
			dictDelete := d.dictDelete(fmt.Sprintf("key%d", i))
			if !dictDelete {
				t.Errorf("dictDelete failed, key: %s", fmt.Sprintf("key%d", i))
			}
		}
	})
	d.dictPrintStats()
	for i := 0; i < 40960; i++ {
		d.dictAdd(fmt.Sprintf("key%d", i), i)
	}
	d.dictPrintStats()

}

func TestDict_dictGetRandomKey(t *testing.T) {
	d := NewDict()
	for i := 0; i < 409600; i++ {
		d.dictAdd(fmt.Sprintf("key%d", i), i)
	}
	d.dictPrintStats()
	t.Run("dictGetRandomKey", func(t *testing.T) {
		for i := 0; i < 409600; i++ {
			dictGetRandomKey := d.dictGetRandomKey()
			if dictGetRandomKey == nil {
				t.Errorf("dictGetRandomKey failed")
			}
		}
	})
	for i := 0; i < 16; i++ {
		fmt.Println(d.dictGetRandomKey().key)
	}
	d.dictPrintStats()
}

func TestDict_keys(t *testing.T) {
	d := NewDict()
	for i := 0; i < 409600; i++ {
		d.dictAdd(fmt.Sprintf("key%d", i), i)
	}
	d.dictPrintStats()
	t.Run("keys", func(t *testing.T) {
		keys := d.keys()
		if len(keys) != 409600 {
			t.Errorf("keys failed")
		}
	})
	d.dictPrintStats()
}

func TestDict_values(t *testing.T) {
	d := NewDict()
	for i := 0; i < 409600; i++ {
		d.dictAdd(fmt.Sprintf("key%d", i), i)
	}
	d.dictPrintStats()
	t.Run("values", func(t *testing.T) {
		values := d.values()
		if len(values) != 409600 {
			t.Errorf("values failed")
		}
	})
	d.dictPrintStats()
}
